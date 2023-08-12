package main

import (
	"catchreview-api-app/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.ConfInitialize()
	if err != nil {
		log.Fatalln("[main] failed config initialize err : ", err)
		return
	}

	router := gin.Default()
	router.GET("/api/health-check")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ApiPort),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}
