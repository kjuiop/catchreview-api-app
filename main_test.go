package main

import (
	"catchreview-api-app/config"
	"catchreview-api-app/internal/delivery/http/handler"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

func TestMainFunction(t *testing.T) {
	gin.SetMode(gin.TestMode)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	cfg := &config.Config{ApiPort: "8088"}

	go main()

	time.Sleep(time.Second)

	client := http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:"+cfg.ApiPort+"/api/health-check", nil)
	resp, err := client.Do(req)
	assert.NoError(t, err)

	quit <- syscall.SIGINT

	time.Sleep(time.Second)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestApiHandler_CloseWithContext(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()

	cfg := &config.Config{ApiPort: "0"}
	ctx, cancel := context.WithCancel(context.Background())

	server := &http.Server{
		Addr:    ":" + cfg.ApiPort,
		Handler: router,
	}

	wg := sync.WaitGroup{}
	// 테스트: 핸들러의 CloseWithContext 호출
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// 1초 후 종료 시그널 전송
		time.Sleep(time.Second * 1)
		quit <- syscall.SIGINT
	}()

	wg.Add(1)
	go closeWithContext(ctx, cancel, server, quit, &wg)

	wg.Wait() // 모든 고루틴이 종료될 때까지 대기

	assert.ErrorIs(t, context.Canceled, ctx.Err())
}

func TestApiHandler_ServeHttpServer(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()
	group := router.Group("/api")

	cfg := &config.Config{ApiPort: "8098"}
	ctx, cancel := context.WithCancel(context.Background())

	handler.NewApiHandler(cfg, group)

	server := &http.Server{
		Addr:    ":" + cfg.ApiPort,
		Handler: router,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go serveHttpServer(ctx, server, &wg)

	// 테스트용 HTTP 클라이언트 생성
	client := http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:"+cfg.ApiPort+"/api/health-check", nil)
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
	go closeWithContext(ctx, cancel, server, quit, &wg)

	wg.Wait() // 모든 고루틴이 종료될 때까지 대기

	assert.ErrorIs(t, context.Canceled, ctx.Err())
}
