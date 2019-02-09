package main

import (
	"Dotato-di-una-libreria/backend/controller"
	"Dotato-di-una-libreria/backend/logger"
	"Dotato-di-una-libreria/backend/middleware"
	"Dotato-di-una-libreria/backend/system"
	"Dotato-di-una-libreria/backend/util"
	"context"
	"fmt"

	"google.golang.org/api/option"

	firebase "firebase.google.com/go"

	"github.com/labstack/echo"

	"cloud.google.com/go/datastore"
	"google.golang.org/appengine"

	"github.com/jinzhu/gorm"
)

func main() {
	// --------------------------------------------------------------
	// ロガーの初期化
	// --------------------------------------------------------------
	lgr := logger.ZapLogger()
	appLgr := logger.NewLogger(lgr.Sugar())

	// --------------------------------------------------------------
	// DatastoreからDB接続情報取得
	// --------------------------------------------------------------
	csqlSetting := system.CloudSQLSetting{}
	if util.IsLocal() {
		// 面倒なのでローカルでは直接入れてしまう
		csqlSetting = system.LocalDatastoreSetup()
	} else {
		ctx := context.Background()
		dsClient, err := datastore.NewClient(ctx, appengine.AppID(ctx))
		if err != nil {
			appLgr.Panicw(err.Error())
		}
		key := datastore.NameKey("CloudSQLSetting", "ID", nil)
		if err := dsClient.Get(ctx, key, &csqlSetting); err != nil {
			appLgr.Errorw("Error at datastore.Get", "Kind", key.Kind, "Name", key.Name, "ID", key.ID)
			appLgr.Panicw(err.Error())
		}
	}

	dataSource := fmt.Sprintf(system.InstanceStrFormat(),
		csqlSetting.User, csqlSetting.Password, csqlSetting.InstanceStr, csqlSetting.DBName)
	if util.IsLocal() {
		appLgr.Infow("-", "dataSource", dataSource)
	}

	// --------------------------------------------------------------
	// Cloud SQLへのコネクション取得
	// --------------------------------------------------------------
	var err error
	db, err := gorm.Open("mysql", dataSource)
	if err != nil {
		appLgr.Panicw(err.Error())
	}
	if db == nil {
		appLgr.Panicw("can not connect to Cloud SQL")
	}
	defer db.Close()

	// http://doc.gorm.io/advanced.html#generic-database-interface-sqldb
	// http://doc.gorm.io/advanced.html#logger
	db.LogMode(true)
	if err := db.DB().Ping(); err != nil {
		appLgr.Panicw("no response from Cloud SQL", "err", err)
	}
	db.DB().SetMaxIdleConns(csqlSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(csqlSetting.MaxOpenConns)

	// TODO: 本来はローカル環境用だけど、ひとまずGAE環境でも同じレコードを投入してしまう
	// ダミーレコードを投入しておく
	err = system.SetupLocalRecord(db)
	if err != nil {
		appLgr.Panicw(err.Error())
	}

	// -----------------------------------------------------
	// Firebase Admin SDK初期化
	// -----------------------------------------------------
	var firebaseApp *firebase.App
	if util.IsLocal() {
		opt := option.WithCredentialsFile("/home/sky0621/work/var/firebase/firebase-adminsdk.json")
		firebaseApp, err = firebase.NewApp(context.Background(), nil, opt)
	} else {
		firebaseApp, err = firebase.NewApp(context.Background(), nil)
	}
	if err != nil {
		appLgr.Panicw("error initializing app", "err", err)
	}

	// --------------------------------------------------------------
	// Webサーバーのセッティング
	// --------------------------------------------------------------
	// https://echo.labstack.com/guide
	e := echo.New()

	// FIXME: DatastoreないしCloud SQLからAPIキーを取得するコード
	apiKey := "DUMMY"

	// 標準ミドルウェアセッティング
	middleware.SetupBasic(e, apiKey)

	// カスタムミドルウェアセッティング
	middleware.SetupCustom(e, appLgr, db, firebaseApp)

	// ルーティング設定の起点
	controller.Routing(e)

	// TODO: Go1.11ではappengine.Main()は必須ではないが、↓の問題があるので、ひとまず使っておく。
	// https://github.com/gcpug/nouhau/issues/71
	appengine.Main()
	//e.Logger.Fatal(e.Start(":8080"))
}
