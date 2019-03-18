package valueobject

const (
	// Publish ... 公開
	Publish PublishFlg = true

	// Private ... 非公開
	Private PublishFlg = false
)

// PublishFlg ... 特定情報の公開設定
type PublishFlg bool
