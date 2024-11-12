package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"uv-chat-api-server-golang/app/controller"
	"uv-chat-api-server-golang/app/repository"
	"uv-chat-api-server-golang/app/router"
	"uv-chat-api-server-golang/app/service"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		repository.Modules,
		service.Modules,
		controller.Modules,
		router.Modules,
	)
	setupSignalHandler(app)
	app.Run()
}

func setupSignalHandler(app *fx.App) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigs
		log.Printf("Received signal: %s. Initiating graceful shutdown...", sig)
		app.Stop(context.Background()) // graceful shutdown 트리거
	}()
}
