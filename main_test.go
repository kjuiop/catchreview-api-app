package main

import (
	"catchreview-api-app/config"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"os/signal"
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
