package repository

import (
	"Dotato-di-una-libreria/backend2/domain/error"
	"Dotato-di-una-libreria/backend2/domain/model"
	vo "Dotato-di-una-libreria/backend2/domain/valueobject"
	"context"
)

// NoticeCommandRepository ... 「お知らせ」データへのCRUDのうち、Commandに該当するCUDを担う。
type NoticeCommandRepository interface {
	/*
	 * 新規登録
	 */
	// Create ... 引数で渡された「お知らせ」データ１件を作成する。
	Create(context.Context, *model.Notice) error.ApplicationError

	// CreateBatch ... 引数で渡された「お知らせ」データ複数件を作成する。
	CreateBatch(context.Context, []*model.Notice) error.ApplicationError

	/*
	 * 更新
	 */
	// UpdateByUniqueID ... 引数で渡された「お知らせ」データに含まれるIDを条件に１件を更新する。
	UpdateByUniqueID(context.Context, *model.Notice) error.ApplicationError

	// UpdateByCondition ... 引数で渡された「お知らせ」データ更新条件に合致する複数の「お知らせ」データを更新する。
	UpdateByCondition(context.Context, *model.NoticeCommandCondition, *model.Notice) error.ApplicationError

	// UpdateBatch ... 引数で渡された「お知らせ」データ複数件を更新する。
	UpdateBatch(context.Context, []*model.Notice) error.ApplicationError

	/*
	 * 削除
	 */
	// DeleteByUniqueID ... 引数で渡されたユニークIDで特定される「お知らせ」データ１件を削除する。
	DeleteByUniqueID(context.Context, vo.UniqueID) error.ApplicationError

	// DeleteByCondition ... 引数で渡された「お知らせ」データ削除条件に合致する複数の「お知らせ」データを削除する。
	DeleteByCondition(context.Context, *model.NoticeCommandCondition) error.ApplicationError

	// DeleteBatch ... 引数で渡されたユニークIDで特定される「お知らせ」データ複数件を削除する。
	DeleteBatch(context.Context, []vo.UniqueID) error.ApplicationError
}

// NoticeQueryRepository ... 「お知らせ」データへのCRUDのうち、Queryに該当するRを担う。
type NoticeQueryRepository interface {
	// GetByUniqueID ... 引数で渡されたユニークIDで特定される「お知らせ」データ１件を取得する。（取得対象が存在しない場合は nil を返却する。）
	GetByUniqueID(context.Context, vo.UniqueID) (*model.Notice, error.ApplicationError)

	// GetByCondition ... 引数で渡された「お知らせ」データ取得条件に合致する複数の「お知らせ」データを返却する。（取得条件が nil の場合は全ての「お知らせ」データを返却する。）
	GetByCondition(context.Context, *model.NoticeReadCondition) ([]*model.Notice, error.ApplicationError)
}
