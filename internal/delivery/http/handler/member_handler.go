package handler

import (
	"catchreview-api-app/domain"
	"github.com/gin-gonic/gin"
)

type MemberHandler struct {
	MUsecase domain.MemberUsecase
}

func NewMemberHandler(group *gin.RouterGroup, us domain.MemberUsecase) {
	handler := &MemberHandler{
		MUsecase: us,
	}

	group.POST("/members", handler.Store)
}

func (m *MemberHandler) Store(g *gin.Context) {
}
