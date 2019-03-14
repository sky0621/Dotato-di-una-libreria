package repository

import (
	"Dotato-di-una-libreria/backend2/domain/model"
	vo "Dotato-di-una-libreria/backend2/domain/valueobject"
	"context"
)

// NoticeCommandRepository ... 「お知らせ」データへのCRUDのうち、Commandに該当するCUDを担う。
type NoticeCommandRepository interface {
	// Create ... 引数で渡された「お知らせ」データ１件を作成する。（作成した「お知らせ」データを返却する。）
	Create(context.Context, *model.Notice) (*model.Notice, error)

	// Update ... 引数で渡された「お知らせ」データ１件を更新する。（更新した「お知らせ」データを返却する。）
	Update(context.Context, *model.Notice) (*model.Notice, error)

	// Delete ... 引数で渡されたユニークIDで特定される「お知らせ」データ１件を削除する。（削除した「お知らせ」データを返却する。）
	Delete(context.Context, vo.UniqueID) (*model.Notice, error)

	// FIXME: 要件に応じて、 CreateBatch(~~) や UpdateBatch(~~) といった、複数の「お知らせ」データを一括で処理するメソッドも定義する。
}

// NoticeQueryRepository ... 「お知らせ」データへのCRUDのうち、Queryに該当するRを担う。
type NoticeQueryRepository interface {
	// Read ... 引数で渡された「お知らせ」データ取得条件に合致する複数の「お知らせ」データを返却する。（取得条件が nil の場合は全ての「お知らせ」データを返却する。）
	Read(context.Context, *model.NoticeReadCondition) ([]*model.Notice, error)

	// FIXME: １件取得用のメソッド定義を検討する。
}
