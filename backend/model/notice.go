package model

import (
	"Dotato-di-una-libreria/backend/logger"

	"github.com/jinzhu/gorm"
)

// REF: http://doc.gorm.io/

// Notice ...
type Notice struct {
	ID       string `json:"id" gorm:"column:id;primary_key"`
	Sentence string `json:"sentence" gorm:"column:sentence;type:varchar(256);not null"`
	AuditItem
}

// TableName ...
func (n *Notice) TableName() string {
	return "notice"
}

// IsDto ...
func (n *Notice) IsDto() bool { return true }

// NoticeDao ...
type NoticeDao interface {
	ListNotice() ([]*Notice, error)
}

type noticeDao struct {
	lgr logger.AppLogger
	db  *gorm.DB
}

// NewNoticeDao ...
func NewNoticeDao(lgr logger.AppLogger, db *gorm.DB) NoticeDao {
	return &noticeDao{
		lgr: lgr,
		db:  db,
	}
}

// ListNotice ...
func (d *noticeDao) ListNotice() ([]*Notice, error) {
	log := d.lgr.Path("model/noticeDao#ListNotice")
	log.Infow("MODEL__Start")

	var res []*Notice
	if err := d.db.Find(&res).Error; err != nil {
		d.lgr.Errorw(err.Error())
		return nil, err
	}

	log.Infow("MODEL__End")
	return res, nil
}
