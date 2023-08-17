package http

import (
	"catchreview-api-app/config"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
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

func (a *ApiHandler) ServeHttpServer(ctx context.Context, srv *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("ServeHttpServer in")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
		return
	}
}

func (a *ApiHandler) CloseWithContext(ctx context.Context, cancel context.CancelFunc, srv *http.Server, quit chan os.Signal, wg *sync.WaitGroup) {
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
