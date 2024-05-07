package db

import (
	"fmt"
	"user-service/config"

	"github.com/go-pg/pg/v10"
)

func NewPostgresDB(cfg *config.DBConfig) (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Database,
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	return db, nil
}
