package handler

import (
	"catchreview-api-app/config"
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

func TestApiHandler_CloseWithContext(t *testing.T) {

	gin.SetMode(gin.TestMode)
	router := gin.New()

	cfg := &config.Config{ApiPort: "8080"}
	ctx, cancel := context.WithCancel(context.Background())
	handler := NewApiHandler(cfg, ctx, cancel)

	server := &http.Server{
		Addr:    ":" + cfg.ApiPort,
		Handler: router,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go handler.ServeHttpServer(server, &wg)

	// 테스트: 핸들러의 CloseWithContext 호출
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// 1초 후 종료 시그널 전송
		time.Sleep(time.Second)
		quit <- syscall.SIGINT
	}()

	wg.Add(1)
	go handler.CloseWithContext(server, quit, &wg)

	wg.Wait() // 모든 고루틴이 종료될 때까지 대기

	assert.ErrorIs(t, ctx.Err(), context.Canceled)
}
