package handler

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
	ctx       context.Context
	ctxCancel context.CancelFunc
	cfg       *config.Config
}

func NewApiHandler(cfg *config.Config, ctx context.Context, cancel context.CancelFunc) *ApiHandler {
	return &ApiHandler{
		cfg:       cfg,
		ctx:       ctx,
		ctxCancel: cancel,
	}
}

func (a *ApiHandler) HealthCheck(gCtx *gin.Context) {
	gCtx.JSON(http.StatusOK, map[string]string{"result": "success"})
}

func (a *ApiHandler) ServeHttpServer(srv *http.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Println("ServeHttpServer in")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
		return
	}
}

func (a *ApiHandler) CloseWithContext(srv *http.Server, quit chan os.Signal, wg *sync.WaitGroup) {

	defer wg.Done()

	for {
		select {
		case <-quit:
			log.Printf("Received exit signal: %v\n", quit)
			a.ctxCancel()
		case <-a.ctx.Done():
			log.Println("Context done, initiating graceful shutdown...")

			if err := srv.Shutdown(a.ctx); err != nil {
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
