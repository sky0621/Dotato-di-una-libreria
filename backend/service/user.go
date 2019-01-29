package service

import (
	"Dotato-di-una-libreria/backend/logger"
	"Dotato-di-una-libreria/backend/middleware"
	"Dotato-di-una-libreria/backend/model"

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
}

// NewUser ...
func NewUser(ctx middleware.CustomContext) UserService {
	return &userService{
		lgr:         ctx.GetLog(),
		db:          ctx.GetDB(),
		firebaseApp: ctx.GetFirebaseApp(),
	}
}

// ListUser ...
func (n *userService) CreateUser(u *model.User) error {
	n.lgr.Path("service/ListUser").Infow("Start")

	// FIXME: トランザクション張って、ここでDB登録後にFirebaseAuthのCreateUserを行う！

	return model.NewUserDao(n.lgr, n.db, n.firebaseApp).CreateUser(u)
}
