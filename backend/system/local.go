package system

import (
	"github.com/jinzhu/gorm"
)

// LocalDatastoreSetup ...
func LocalDatastoreSetup() CloudSQLSetting {
	dbc := &CloudSQLSetting{
		ID:          1,
		User:        "testuser",
		Password:    "testpass",
		InstanceStr: "127.0.0.1:3306",
		DBName:      "testdb"}
	return *dbc
}

// SetupLocalRecord ...
func SetupLocalRecord(db *gorm.DB) error {
	db.Exec("INSERT INTO notice (id, sentence, create_user, created_at) VALUES('1', '１つめのお知らせです。', 'user', CURRENT_TIMESTAMP)")
	db.Exec("INSERT INTO notice (id, sentence, create_user, created_at) VALUES('2', '２つめのお知らせです。', 'user', CURRENT_TIMESTAMP)")
	db.Exec("INSERT INTO notice (id, sentence, create_user, created_at) VALUES('3', '３つめのお知らせです。', 'user', CURRENT_TIMESTAMP)")
	db.Exec("INSERT INTO notice (id, sentence, create_user, created_at) VALUES('4', '４つめのお知らせです。', 'user', CURRENT_TIMESTAMP)")

	return nil
}
