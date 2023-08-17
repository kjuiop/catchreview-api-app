package handler

import (
	"catchreview-api-app/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiHandler_HealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	cfg := &config.Config{}
	handler := NewApiHandler(cfg)
	router.GET("/health", handler.HealthCheck)

	// 테스트용 HTTP 요청 생성
	req := httptest.NewRequest("GET", "/health", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	// 테스트 결과 확인
	assert.Equal(t, http.StatusOK, resp.Code)

	expectedBody := `{"result":"success"}`
	assert.Equal(t, expectedBody, resp.Body.String())
}
