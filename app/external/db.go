package external

import (
	"fmt"
	"uv-chat-api-server-golang/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func mustDB(config *config.Config) *gorm.DB {
	var db *gorm.DB
	var err error
	switch config.Database.DBType {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.Database.DSN), &gorm.Config{})
	case "mysql":
		db, err = gorm.Open(mysql.Open(config.Database.DSN), &gorm.Config{})
	default:
		err = fmt.Errorf("unsupported database type: %s", config.Database.DBType)
	}
	if err != nil {
		panic(err)
	}
	return db
}
