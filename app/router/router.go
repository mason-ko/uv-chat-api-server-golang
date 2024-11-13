package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"uv-chat-api-server-golang/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(
		newRouter,
		newServer,
	),
	fx.Invoke(func(*http.Server) {}),
)

func newRouter(controller domain.Controller) *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})
	return router
}

func newServer(lc fx.Lifecycle, router *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			// Graceful shutdown with a timeout
			ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
			defer cancel()
			fmt.Println("Shutting down server gracefully...")
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
