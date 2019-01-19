package service

import (
	"Dotato-di-una-libreria/backend/logger"
	"Dotato-di-una-libreria/backend/middleware"
	"Dotato-di-una-libreria/backend/model"

	"github.com/jinzhu/gorm"
)

// Notice ...
type Notice interface {
	ListNotice() ([]*model.Notice, error)
}

type noticeService struct {
	lgr logger.AppLogger
	db  *gorm.DB
}

// NewNoticeService ...
func NewNoticeService(ctx middleware.CustomContext) Notice {
	return &noticeService{
		lgr: ctx.GetLog(),
		db:  ctx.GetDB(),
	}
}

// ListNotice ...
func (n *noticeService) ListNotice() ([]*model.Notice, error) {
	n.lgr.Path("service/ListNotice").Infow("Start")
	return model.NewNoticeDao(n.lgr, n.db).ListNotice()
}
