package model

import (
	"Dotato-di-una-libreria/backend/logger"

	firebase "firebase.google.com/go"

	"github.com/jinzhu/gorm"
)

// REF: http://doc.gorm.io/

// User ...
type User struct {
	ID       string `gorm:"column:id;primary_key"`
	Name     string `gorm:"column:name;type:varchar(256);not null"`
	Mail     string `gorm:"column:mail;type:varchar(256);not null"`
	Password string `gorm:"-"`
	AuditItem
}

// TableName ...
func (n *User) TableName() string {
	return "user"
}

// IsDto ...
func (n *User) IsDto() bool { return true }

// UserDao ...
type UserDao interface {
	CreateUser(u *User) error
}

type userDao struct {
	lgr         logger.AppLogger
	db          *gorm.DB
	firebaseApp *firebase.App
}

// NewUserDao ...
func NewUserDao(lgr logger.AppLogger, db *gorm.DB, firebaseApp *firebase.App) UserDao {
	return &userDao{
		lgr:         lgr,
		db:          db,
		firebaseApp: firebaseApp,
	}
}

// CreateUser ...
func (d *userDao) CreateUser(u *User) error {
	log := d.lgr.Path("model/UserDao#CreateUser")
	log.Infow("MODEL__Start")

	if err := d.db.Create(&u).Error; err != nil {
		d.lgr.Errorw(err.Error())
		return err
	}

	log.Infow("MODEL__End")
	return nil
}
