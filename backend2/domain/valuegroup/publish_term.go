package valuegroup

import "time"

// PublishTerm ... 公開期間
type PublishTerm struct {
	// PublishedFrom ... お知らせの公開開始日時（ nil の場合は即時開始）
	PublishedFrom *time.Time
	// PublishedTo ... お知らせの公開終了日時（ nil の場合は終了）
	PublishedTo *time.Time
}
