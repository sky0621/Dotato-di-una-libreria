package valueobject

const (
	// NoticePublish ... 「お知らせ」公開
	NoticePublish NoticePublishFlg = true

	// NoticePrivate ... 「お知らせ」非公開
	NoticePrivate NoticePublishFlg = false
)

// NoticePublishFlg ... 「お知らせ」の公開設定
type NoticePublishFlg bool
