package storage

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/zsandibe/effective_mobile_task/config"
	"github.com/zsandibe/effective_mobile_task/pkg"
)

func NewPostgres(config config.Config) (*sqlx.DB, error) {
	dbSource := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", config.Database.Driver, config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DatabaseName)
	db, err := sqlx.Open(config.Database.Driver, dbSource)
	if err != nil {
		pkg.ErrorLog.Printf("Error opening database: %v", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		pkg.ErrorLog.Printf("Error connecting database: %v", err)
		return nil, err
	}
	return db, nil
}
