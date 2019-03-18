package aggregate

import (
	vo "Dotato-di-una-libreria/backend2/domain/valueobject"
)

// Publish ...
func Publish(term *PublishTerm) *PublishControl {
	return &PublishControl{
		isPublish: vo.Publish,
		term:      term,
	}
}

// Private ...
func Private() *PublishControl {
	return &PublishControl{
		isPublish: vo.Private,
	}
}

// PublishControl ... 特定情報の公開制御
type PublishControl struct {
	// IsPublish ... 公開フラグ
	isPublish vo.PublishFlg
	// Term ... 公開期間
	term *PublishTerm
}
