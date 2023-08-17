package main

import (
	"catchreview-api-app/config"
	"catchreview-api-app/handler"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}

	quit := make(chan os.Signal, 1)
	defer close(quit)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	cfg, err := config.ConfInitialize()
	if err != nil {
		log.Fatalln("[main] failed config initialize err : ", err)
		return
	}

	a := handler.NewApiHandler(cfg, ctx, cancel)

	router := gin.Default()
	router.GET("/api/health-check", a.HealthCheck)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ApiPort),
		Handler: router,
	}

	wg.Add(1)
	go a.CloseWithContext(srv, quit, &wg)

	wg.Add(1)
	go a.ServeHttpServer(srv, &wg)

	wg.Wait()
}
