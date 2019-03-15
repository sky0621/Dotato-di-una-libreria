package model

import (
	vo "Dotato-di-una-libreria/backend2/domain/valueobject"
)

// Notice ... 「お知らせ」データ定義
type Notice struct {
	ID       vo.UniqueID
	Title    string
	Sentence string
	Severity vo.NoticeSeverity
}

// NoticeCommandCondition ... 条件に該当する「お知らせ」データを決定するために利用
type NoticeCommandCondition struct {
	// FIXME 複数の Notice を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

// NoticeReadCondition ... 条件に該当する「お知らせ」データを決定するために利用
type NoticeReadCondition struct {
	// FIXME 複数の Notice を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}
