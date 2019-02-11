package service

import (
	"Dotato-di-una-libreria/backend/logger"
	"Dotato-di-una-libreria/backend/middleware"
	"Dotato-di-una-libreria/backend/model"
	"Dotato-di-una-libreria/backend/util"
	"context"

	"firebase.google.com/go/auth"

	firebase "firebase.google.com/go"

	"github.com/jinzhu/gorm"
)

// UserService ...
type UserService interface {
	CreateUser(u *model.User) error
}

type userService struct {
	lgr         logger.AppLogger
	db          *gorm.DB
	firebaseApp *firebase.App
	requestCtx  context.Context
}

// NewUser ...
func NewUser(ctx middleware.CustomContext, requestCtx context.Context) UserService {
	return &userService{
		lgr:         ctx.GetLog(),
		db:          ctx.GetDB(),
		firebaseApp: ctx.GetFirebaseApp(),
		requestCtx:  requestCtx,
	}
}

// ListUser ...
func (s *userService) CreateUser(u *model.User) error {
	s.lgr.Path("service/ListUser").Infow("Start")

	tx := s.db.Begin()
	defer func() {
		if tx != nil {
			db := tx.Commit()
			if err := db.Error; err != nil {
				s.lgr.Errorw("Transaction commit failed.", "error", err)
			}
		}
	}()

	u.ID = util.CreateUniqueID()
	err := model.NewUserDao(s.lgr, tx, s.firebaseApp).CreateUser(u)
	if err != nil {
		s.lgr.Errorw("@userDao#CreateUser", "error", err)
		return err
	}

	fbAuth, err := s.firebaseApp.Auth(s.requestCtx)
	if err != nil {
		s.lgr.Errorw("@firebase.GetAuth", "error", err)
		return err
	}
	fbUser := &auth.UserToCreate{}
	fbUser.Email(u.Mail)
	fbUser.Password(u.Password)
	_, err = fbAuth.CreateUser(s.requestCtx, fbUser)
	if err != nil {
		s.lgr.Errorw("@firebase.CreateUser", "error", err)
		return err
	}

	return nil
}
