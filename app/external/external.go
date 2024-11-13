package external

import (
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/internal/config"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Modules = fx.Options(
	fx.Provide(newExternal),
)

type external struct {
	db *gorm.DB
}

// DB implements domain.External.
func (e *external) DB() *gorm.DB {
	return e.db
}

func newExternal(config *config.Config) domain.External {
	return &external{db: mustDB(config)}
}
