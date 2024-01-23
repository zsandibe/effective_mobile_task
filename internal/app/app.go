package app

import (
	"github.com/zsandibe/effective_mobile_task/config"
	"github.com/zsandibe/effective_mobile_task/internal/repository"
	"github.com/zsandibe/effective_mobile_task/internal/service"
	"github.com/zsandibe/effective_mobile_task/internal/storage"
	"github.com/zsandibe/effective_mobile_task/pkg"
)

func Start() {
	config, err := config.NewConfig()
	if err != nil {
		pkg.ErrorLog.Printf("Problems with loading config: %v", err)
		return
	}
	pkg.InfoLog.Println("Config loaded")

	db, err := storage.NewPostgres(*config)
	if err != nil {
		pkg.ErrorLog.Printf("Can`t connect to database: %v", err)
		return
	}
	pkg.InfoLog.Println("Database loaded")

	repository := repository.NewRepository(db)
	pkg.InfoLog.Println("Repository loaded")

	service := service.NewPersonService(repository, config)
	pkg.InfoLog.Println("Service loaded")

}
