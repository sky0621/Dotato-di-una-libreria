package valuegroup

import (
	vo "Dotato-di-una-libreria/backend2/domain/valueobject"
)

// PublishControl ... 特定情報の公開制御
type PublishControl struct {
	// IsPublish ... 公開フラグ
	IsPublish vo.PublishFlg
	// Term ... 公開期間
	Term PublishTerm
}
