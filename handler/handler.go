package handler

import (
	"catchreview-api-app/config"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiHandler struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	cfg       *config.Config
}

func NewApiHandler(cfg *config.Config) *ApiHandler {
	return &ApiHandler{
		cfg: cfg,
	}
}

func (a *ApiHandler) HealthCheck(gCtx *gin.Context) {
	gCtx.JSON(http.StatusOK, map[string]string{"result": "success"})
}
