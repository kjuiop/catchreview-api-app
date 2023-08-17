package http

import (
	"catchreview-api-app/config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
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

func TestApiHandler_ServeHttpServer(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()

	cfg := &config.Config{ApiPort: "8098"}
	ctx, cancel := context.WithCancel(context.Background())
	handler := NewApiHandler(cfg)
	router.GET("/health", handler.HealthCheck)

	server := &http.Server{
		Addr:    ":" + cfg.ApiPort,
		Handler: router,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go handler.ServeHttpServer(ctx, server, &wg)

	// 테스트용 HTTP 클라이언트 생성
	client := http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:"+cfg.ApiPort+"/health", nil)
	resp, err := client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// 1초 후 종료 시그널 전송
		time.Sleep(time.Second)
		quit <- syscall.SIGINT
	}()

	wg.Add(1)
	go handler.CloseWithContext(ctx, cancel, server, quit, &wg)

	wg.Wait() // 모든 고루틴이 종료될 때까지 대기

	assert.ErrorIs(t, context.Canceled, ctx.Err())
}

func TestApiHandler_CloseWithContext(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()

	cfg := &config.Config{ApiPort: "0"}
	ctx, cancel := context.WithCancel(context.Background())
	handler := NewApiHandler(cfg)

	server := &http.Server{
		Addr:    ":" + cfg.ApiPort,
		Handler: router,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go handler.ServeHttpServer(ctx, server, &wg)

	// 테스트: 핸들러의 CloseWithContext 호출
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// 1초 후 종료 시그널 전송
		time.Sleep(time.Second)
		quit <- syscall.SIGINT
	}()

	wg.Add(1)
	go handler.CloseWithContext(ctx, cancel, server, quit, &wg)

	wg.Wait() // 모든 고루틴이 종료될 때까지 대기

	assert.ErrorIs(t, context.Canceled, ctx.Err())
}
