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

	const noticeQueryBase = "SELECT COUNT(1) AS count FROM notice WHERE id = "
	const noticeExecBase = "INSERT INTO notice (id, sentence, create_user, created_at) VALUES"
	const userQueryBase = "SELECT COUNT(1) AS count FROM user WHERE id = "
	const userExecBase = "INSERT INTO user (id, name, mail, create_user, created_at) VALUES"
	sqls := [][]string{
		[]string{
			noticeQueryBase + "1",
			noticeExecBase + "('1', '１つめのお知らせです。', 'user', CURRENT_TIMESTAMP)",
		},
		[]string{
			noticeQueryBase + "2",
			noticeExecBase + "('2', '２つめのお知らせです。', 'user', CURRENT_TIMESTAMP)",
		},
		[]string{
			noticeQueryBase + "3",
			noticeExecBase + "('3', '３つめのお知らせです。', 'user', CURRENT_TIMESTAMP)",
		},
		[]string{
			noticeQueryBase + "4",
			noticeExecBase + "('4', '４つめのお知らせです。', 'user', CURRENT_TIMESTAMP)",
		},
		[]string{
			userQueryBase + "1",
			userExecBase + "('1', 'ユーザー１', 'user01@example.com', 'user', CURRENT_TIMESTAMP)",
		},
	}

	for _, sql := range sqls {
		n := &count{}
		if err := tx.Raw(sql[0]).Scan(n).Error; err != nil {
			tx.Rollback()
			return err
		}
		if n != nil && n.Count > 0 {
			continue
		}
		if err := tx.Exec(sql[1]).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}

type count struct {
	Count int `gorm:"column:count"`
}
