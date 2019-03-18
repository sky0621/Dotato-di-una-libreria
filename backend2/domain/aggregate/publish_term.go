package aggregate

import "time"

// NewPublishTerm ...
func NewPublishTerm(from, to *time.Time) *PublishTerm {
	// FIXME: from, to に対するバリデーション（from < to）
	return &PublishTerm{from: from, to: to}
}

// PublishTerm ... 公開期間
type PublishTerm struct {
	// 公開開始日時（ nil の場合は即時開始）
	from *time.Time
	// 公開終了日時（ nil の場合は終了）
	to *time.Time
}

// PublishFrom ... 公開開始日時を返却
func (t *PublishTerm) PublishFrom() *time.Time {
	if t == nil {
		return nil
	}
	return t.from
}

// PublishTo ... 公開終了日時を返却
func (t *PublishTerm) PublishTo() *time.Time {
	if t == nil {
		return nil
	}
	return t.to
}
