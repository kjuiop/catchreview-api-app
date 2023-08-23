package handler

import (
	"catchreview-api-app/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ApiHandler struct {
	cfg *config.Config
}

func NewApiHandler(cfg *config.Config, group *gin.RouterGroup) {
	handler := &ApiHandler{
		cfg: cfg,
	}
	group.GET("/health-check", handler.HealthCheck)
}

func (a *ApiHandler) HealthCheck(gCtx *gin.Context) {
	gCtx.JSON(http.StatusOK, map[string]string{"result": "success"})
}
