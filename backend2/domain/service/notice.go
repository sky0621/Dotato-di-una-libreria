package service

import (
	"Dotato-di-una-libreria/backend2/domain/error"
	"Dotato-di-una-libreria/backend2/domain/model"
	"Dotato-di-una-libreria/backend2/domain/repository"
	"Dotato-di-una-libreria/backend2/domain/service/command"
	"Dotato-di-una-libreria/backend2/domain/service/query"
	"context"
)

// NoticeCommandService ... 「お知らせ」データへのCRUDのうち、Commandに該当するCUDを担う。
type NoticeService interface {
	// Create ... 引数で渡された「お知らせ」データを作成する。
	Create(context.Context, *command.NoticeCommand) error.ApplicationError

	// Update ... 引数で渡された「お知らせ」データを更新する。
	Update(context.Context, *command.NoticeCommand) error.ApplicationError

	// Delete ... 引数で渡された「お知らせ」データを削除する。
	Delete(context.Context, *command.NoticeCommand) error.ApplicationError

	// Get ... 引数で渡された「お知らせ」データを取得する。（取得対象が存在しない場合は nil を返却する。）
	Get(context.Context, *query.NoticeQuery) (*model.Notice, error.ApplicationError)
}

// NewNoticeService ...
func NewNoticeService(
	noticeCommandRepository repository.NoticeCommandRepository,
	noticeQueryRepository repository.NoticeQueryRepository,
) NoticeService {
	if noticeCommandRepository == nil {
		return nil
	}
	if noticeQueryRepository == nil {
		return nil
	}
	return &noticeService{
		noticeCommandRepository: noticeCommandRepository,
		noticeQueryRepository:   noticeQueryRepository,
	}
}

type noticeService struct {
	noticeCommandRepository repository.NoticeCommandRepository
	noticeQueryRepository   repository.NoticeQueryRepository
}

// Create ... 引数で渡された「お知らせ」データを作成する。
func (r *noticeService) Create(ctx context.Context, cmd *command.NoticeCommand) error.ApplicationError {

	return r.noticeCommandRepository.Create(ctx)
}

// Update ... 引数で渡された「お知らせ」データを更新する。
func (r *noticeService) Update(ctx context.Context, cmd *command.NoticeCommand) error.ApplicationError {

}

// Delete ... 引数で渡された「お知らせ」データを削除する。
func (r *noticeService) Delete(ctx context.Context, cmd *command.NoticeCommand) error.ApplicationError {

}

// Get ... 引数で渡された「お知らせ」データを取得する。（取得対象が存在しない場合は nil を返却する。）
func (r *noticeService) Get(ctx context.Context, q *query.NoticeQuery) (*model.Notice, error.ApplicationError) {

}
