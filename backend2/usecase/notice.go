package usecase

import (
	"Dotato-di-una-libreria/backend2/domain/error"
	"Dotato-di-una-libreria/backend2/domain/model"
	"Dotato-di-una-libreria/backend2/domain/repository"
	"Dotato-di-una-libreria/backend2/usecase/input"
	"context"
)

// NewNoticeUsecase ...
func NewNoticeUsecase(commandRepository repository.NoticeCommandRepository, queryRepository repository.NoticeQueryRepository) NoticeUsecase {
	return &noticeUsecase{}
}

// NoticeUsecase ... 「お知らせ」に対するユースケース定義
type NoticeUsecase interface {
	AddNotice(context.Context, *input.Notice) error.ApplicationError
}

type noticeUsecase struct {
	commandRepository repository.NoticeCommandRepository
	queryRepository   repository.NoticeQueryRepository
}

// AddNotice ...
func (u *noticeUsecase) AddNotice(ctx context.Context, input *input.Notice) error.ApplicationError {
	var m *model.Notice
	// FIXME: input -> model.Notice へのコンバーターを実装！
	return u.commandRepository.Create(ctx, m)
}
