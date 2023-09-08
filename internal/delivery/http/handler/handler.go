package handler

import (
	"catchreview-api-app/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// Health Check godoc
// @Summary Api Application Health Check
// @Description Api Application Health Check 를 위한 API 입니다.
// @Accept json
// @Produce json
// @tags HealthCheck
// @Router /health-check [GET]
// @Success 200 {string} result:success
// @Failure 400 "Invalid parameters"
// @Failure 401 "Validation authorization"
// @Failure 500 "health check controller internal exception"
func (a *ApiHandler) HealthCheck(gCtx *gin.Context) {
	gCtx.JSON(http.StatusOK, map[string]string{"result": "success"})
}
