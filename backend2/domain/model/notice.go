package model

import (
	vg "Dotato-di-una-libreria/backend2/domain/valuegroup"
	vo "Dotato-di-una-libreria/backend2/domain/valueobject"
)

// Notice ... 「お知らせ」データ定義
type Notice struct {
	// ID ... お知らせをユニークに特定するID
	ID vo.UniqueID
	// Title ... お知らせの概要を示すタイトル
	Title string
	// Detail ... お知らせの詳細
	Detail string
	// Severity ... お知らせの重要度
	Severity vo.NoticeSeverity
	// PublishControl ... お知らせの公開設定
	PublishControl vg.PublishControl
}

// NoticeCommandCondition ... 条件に該当する「お知らせ」データを決定するために利用
type NoticeCommandCondition struct {
	// FIXME 複数の Notice を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}

// NoticeReadCondition ... 条件に該当する「お知らせ」データを決定するために利用
type NoticeReadCondition struct {
	// FIXME 複数の Notice を持つなり、ページング情報を持つなり、要件にあった取得条件を定義！
}
