package handler

import (
	"catchreview-api-app/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
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
	var member domain.Member

	if err := g.Bind(&member); err != nil {
		g.JSON(http.StatusUnprocessableEntity, ResponseError{
			Message: err.Error(),
		})
		return
	}

	if err := isRequestValid(&member); err != nil {
		g.JSON(http.StatusBadRequest, ResponseError{
			Message: err.Error(),
		})
		return
	}

	ctx := g.Request.Context()
	if err := m.MUsecase.Store(ctx, &member); err != nil {
		g.JSON(http.StatusInternalServerError, ResponseError{
			Message: err.Error(),
		})
		return
	}

	g.JSON(http.StatusOK, member)
	return
}

func isRequestValid(m *domain.Member) error {
	validate := validator.New()
	if err := validate.Struct(m); err != nil {
		return err
	}
	return nil
}
