package usecase

import (
	"catchreview-api-app/domain"
	"context"
)

type memberUsecase struct {
}

func NewMemberUsecase() domain.MemberUsecase {
	return &memberUsecase{}
}

func (m memberUsecase) Store(ctx context.Context, member *domain.Member) {
	panic("implement me")
}
