package gorm

import (
	"fmt"
	"poly_news/internal/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(databaseConfig config.DatabaseConfig, opts ...gorm.Option) (*gorm.DB, error) {
	dialector, err := getDialector(databaseConfig)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(dialector, opts...)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDialector(dbConfig config.DatabaseConfig) (gorm.Dialector, error) {
	switch dbConfig.Type {
	case config.DatabaseTypeSQLite:
		return sqlite.Open(dbConfig.DSN()), nil
	case config.DatabaseTypeMySQL:
		return mysql.Open(dbConfig.DSN()), nil
	default:
		return nil, fmt.Errorf("invalid database type: %s", dbConfig.Type)
	}

}
