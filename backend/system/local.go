package system

import (
	"github.com/jinzhu/gorm"
)

// LocalDatastoreSetup ...
func LocalDatastoreSetup() CloudSQLSetting {
	dbc := &CloudSQLSetting{
		ID:          1,
		User:        "localuser",
		Password:    "localpass",
		InstanceStr: "127.0.0.1:3306",
		DBName:      "localdb"}
	return *dbc
}

// SetupLocalRecord ...
func SetupLocalRecord(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if tx == nil {
			return
		}
		tx.Commit()
	}()
	if tx == nil {
		panic("can not create a db transaction")
	}

	const noticeBase = "INSERT INTO notice (id, sentence, create_user, created_at) VALUES"
	const userBase = "INSERT INTO user (id, name, mail, create_user, created_at) VALUES"
	sqls := []string{
		noticeBase + "('1', '１つめのお知らせです。', 'user', CURRENT_TIMESTAMP)",
		noticeBase + "('2', '２つめのお知らせです。', 'user', CURRENT_TIMESTAMP)",
		noticeBase + "('3', '３つめのお知らせです。', 'user', CURRENT_TIMESTAMP)",
		noticeBase + "('4', '４つめのお知らせです。', 'user', CURRENT_TIMESTAMP)",
		userBase + "('1', 'ユーザー１', 'user01@example.com', 'user', CURRENT_TIMESTAMP)",
	}

	for _, sql := range sqls {
		if err := tx.Exec(sql).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
