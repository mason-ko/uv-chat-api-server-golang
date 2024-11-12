package main

import (
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
	app.Run()
}
