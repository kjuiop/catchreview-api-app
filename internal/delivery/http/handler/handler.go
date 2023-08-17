package handler

import (
	"catchreview-api-app/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiHandler struct {
	cfg *config.Config
}

func NewApiHandler(cfg *config.Config) *ApiHandler {
	return &ApiHandler{
		cfg: cfg,
	}
}

func (a *ApiHandler) HealthCheck(gCtx *gin.Context) {
	gCtx.JSON(http.StatusOK, map[string]string{"result": "success"})
}
