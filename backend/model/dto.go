package model

import (
	"time"
)

// Dto ... 用途はマーカーインタフェースだがダックタイプのためダミーメソッドを定義
type Dto interface {
	IsDto() bool
}

// AuditItem ... 監査項目
type AuditItem struct {
	CreateUser string    `json:"-" gorm:"column:create_user;type:varchar(256);not null"`
	CreatedAt  time.Time `json:"-" gorm:"column:created_at;not null"`
	UpdateUser string    `json:"-" gorm:"column:update_user;type:varchar(256)"`
	UpdatedAt  time.Time `json:"-" gorm:"column:updated_at"`
	// 論理削除に用いる。
	// [参照]http://doc.gorm.io/crud.html#delete
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}

// ID ...
type ID struct {
	ID string `json:"id"`
}
