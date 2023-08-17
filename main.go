package main

import (
	"catchreview-api-app/config"
	"catchreview-api-app/internal/delivery/http/handler"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
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

	h := handler.NewApiHandler(cfg)

	router := gin.Default()
	router.GET("/api/health-check", h.HealthCheck)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ApiPort),
		Handler: router,
	}

	wg.Add(1)
	go closeWithContext(ctx, cancel, srv, quit, &wg)

	wg.Add(1)
	go serveHttpServer(ctx, srv, &wg)

	wg.Wait()
}

func closeWithContext(ctx context.Context, cancel context.CancelFunc, srv *http.Server, quit chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-quit:
			log.Printf("Received exit signal: %v\n", quit)
			cancel()
		case <-ctx.Done():
			log.Println("Context done, initiating graceful shutdown...")

			if err := srv.Shutdown(ctx); err != nil {
				log.Println("Server shutdown error:", err)
				return
			}
			log.Println("Server gracefully stopped")
			return
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

func serveHttpServer(ctx context.Context, srv *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("ServeHttpServer in")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
		return
	}
}
