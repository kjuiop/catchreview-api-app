package usecase

import (
	"catchreview-api-app/domain"
	"context"
	"time"
)

type memberUsecase struct {
	memberRepo     domain.MemberRepository
	contextTimeout time.Duration
}

func NewMemberUsecase(repo domain.MemberRepository, timeout time.Duration) domain.MemberUsecase {
	return &memberUsecase{
		memberRepo:     repo,
		contextTimeout: timeout,
	}
}

func (mu memberUsecase) Store(c context.Context, m *domain.Member) error {
	_, cancel := context.WithTimeout(c, mu.contextTimeout)
	defer cancel()

	currentTime := time.Now()
	m.SignUpBuild(currentTime)

	if err := mu.memberRepo.Store(c, m); err != nil {
		return err
	}
	return nil
}
