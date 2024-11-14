package repository

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"testing"
	"uv-chat-api-server-golang/app/external"
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/internal/config"
)

func Test_baseRepository(t *testing.T) {
	app := fx.New(
		Modules,
		external.Modules,
		config.Modules,

		fx.Provide(func() *testing.T {
			return t
		}),
		fx.Invoke(testBaseRepository),
	)
	ctx := context.Background()
	app.Start(ctx)
	app.Stop(ctx)
}

func testBaseRepository(t *testing.T, repository domain.Repository) {
	fmt.Println("WxWASDSAD")
	t.Run("Create Data", func(t *testing.T) {
		repository.MessageRepository().Create(message.DBMessage{
			ChannelID:         123,
			UserID:            123,
			Content:           "test",
			TranslatedContent: "",
		})
	})

}
