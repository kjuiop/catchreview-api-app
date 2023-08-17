package handler

import (
	"catchreview-api-app/domain"
	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	MUsecase domain.MemberUsecase
}

func NewMemberHandler(us domain.MemberUsecase) *MemberHandler {
	return &MemberHandler{
		MUsecase: us,
	}
}

func (m *MemberHandler) Store(g *gin.Context) {
}
