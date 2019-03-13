package valueobject

const (
	// NoticeSeverityInfo ... 通常レベルのお知らせ
	NoticeSeverityInfo NoticeSeverity = iota + 1

	// NoticeSeverityImportant ... 重要レベルのお知らせ
	NoticeSeverityImportant
)

// NoticeSeverity ... お知らせの重要度
type NoticeSeverity uint
