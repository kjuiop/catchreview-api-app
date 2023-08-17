package domain

import (
	"context"
	"time"
)

type Member struct {
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	Nickname        string    `json:"nickname"`
	PrivacyAgreedAt time.Time `json:"privacy_agreed_at"`
	PolicyAgreedAt  time.Time `json:"policy_agreed_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type MemberUsecase interface {
	Store(context.Context, *Member)
}
