package usecase

import (
	"catchreview-api-app/domain"
	"context"
)

type memberUsecase struct {
	memberRepo domain.MemberRepository
}

func NewMemberUsecase(repo domain.MemberRepository) domain.MemberUsecase {
	return &memberUsecase{
		memberRepo: repo,
	}
}

func (m memberUsecase) Store(ctx context.Context, member *domain.Member) {
	panic("implement me")
}
