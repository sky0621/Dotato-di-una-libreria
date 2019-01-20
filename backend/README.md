# Dotato-di-una-libreria/backend

## 主な使用ライブラリ

- [Echo(Webフレームワーク)](https://echo.labstack.com/guide)
- [Gorm(O/Rマッパー)](http://doc.gorm.io/)
- [Goon(Datastoreアクセッサ)](https://godoc.org/github.com/mjibson/goon) ※使用予定
- [validator.v9(バリデーター)](https://github.com/go-playground/validator/blob/v9/README.md)
- [zap(ロガー)](https://github.com/uber-go/zap)

## goモジュール初期化

「Dotato-di-una-libreria」ディレクトリ直下で実行

**※プロジェクト作成時のみなので各自の実行は不要**

```
$ go mod init Dotato-di-una-libreria/backend
```

## ローカル開発手引き

### ■事前準備

#### (1)dockerインストール

#### (2)docker-composeのインストール

https://docs.docker.com/compose/install/

#### (3)MySQLイメージ取得

```
$ sudo docker pull mysql:5.7.24
```

### ■MySQL起動（docker-composeインストール済み前提）

```
$ sudo docker-compose up
```

※デーモン起動したいなら末尾に「`-d`」

**CRUD操作後のテーブルの中身を確認する際は下記を元にMySQLコンテナに接続**

| Name | Value     |
| :--- | :---      |
| Host | 127.0.0.1 |
| Port | 3306      |
| DB   | testdb    |
| User | testuser  |
| Pass | testpass  |

### ■Datastoreエミュレーター起動（Datastoreの中身を見る必要ないなら不要）

下記参照

https://hackmd.io/s/BJ3d_7q0Q

### ■ローカルサーバ起動

```
$ cd backend/
$ dev_appserver.py app.yaml
```

### ■ローカルでMySQLコンテナ内の状況確認

```
$ docker-compose ps
   Name                 Command             State                 Ports              
-------------------------------------------------------------------------------------
ca-toa_db_1   docker-entrypoint.sh mysqld   Up      0.0.0.0:3306->3306/tcp, 33060/tcp
$
$ docker exec -it ca-toa_db_1 bash
root@cf6ff9d96432:/# mysql -h127.0.0.1 -utestuser -ptestpass
　　〜〜省略〜〜
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| testdb             |
+--------------------+

mysql> use testdb;
　　〜〜省略〜〜
mysql> show tables;
+------------------+
| Tables_in_testdb |
+------------------+
| notice           |
+------------------+
```

## GAE環境デプロイ手引き

※事前に「`$ gcloud config list`」でデプロイ先が正しいことを確認すること！

※クライアント側のビルド結果を「`view/dist`」に出力していないとビューが最新化されないので注意！

```
$ gcloud app deploy
$ gcloud app browse
```

### ■DB接続情報のセッティング

Cloud Datastoreに下記エンティティが必要

| Key | Value |
| :--- | :--- |
| Namespace | [デフォルト] |
| Kind | CloudSQLSetting |
| Key(Nameキー) | ID |

プロパティ

| 名前 | 説明 | タイプ     | 値 |
| :--- | :--- | :--- | :--- |
| ID | Nameキー | 整数 | 1 |
| User | DB接続ユーザ | 文字列 | **[*環境による]** |
| Password | DB接続パスワード | 文字列 | **[*環境による]** |
| InstanceStr | インスタンス接続文字列 | 文字列 | **[*環境による]** |
| DBName | DB名 | 文字列 | **[*環境による]** |
| MaxIdleConns | 最大アイドルコネクション数 | 整数 | **[*環境による]** |
| MaxOpenConns | 最大オープンコネクション数 | 整数 | **[*環境による]** |

### ■APIキーのセッティング

Cloud Datastoreに下記エンティティが必要

| Key | Value |
| :--- | :--- |
| Namespace | [デフォルト] |
| Kind | APIToken |
| Key(Nameキー) | ID |

プロパティ

| 名前 | 説明 | タイプ     | 値 |
| :--- | :--- | :--- | :--- |
| ID | Nameキー | 整数 | 1 |
| Key | APIキー | 文字列 | **[*環境による]** |

## API仕様

### URLパス

ベースパス：「 `/api/v1` 」

例：「お知らせ」機能

[GET] `https://【GCPプロジェクトID】/api/v1/notices`
