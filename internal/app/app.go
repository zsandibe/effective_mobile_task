package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/zsandibe/effective_mobile_task/config"
	"github.com/zsandibe/effective_mobile_task/internal/delivery"
	"github.com/zsandibe/effective_mobile_task/internal/repository"
	"github.com/zsandibe/effective_mobile_task/internal/server"
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
	pkg.InfoLog.Println(config.Api.AgeURL, "laall")
	db, err := storage.NewPostgres(config)
	if err != nil {
		pkg.ErrorLog.Printf("Can`t connect to database: %v", err)
		return
	}
	pkg.InfoLog.Println("Database loaded")

	repository := repository.NewRepository(db)
	pkg.InfoLog.Println("Repository loaded")

	service := service.NewService(repository, config)
	pkg.InfoLog.Println("Service loaded")

	delivery := delivery.NewHandler(service)
	pkg.InfoLog.Println("Delivery loaded")

	server := server.NewServer(*config, delivery.Routes())

	go func() {
		if err := server.Run(); err != nil {
			pkg.ErrorLog.Printf("failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		pkg.ErrorLog.Printf("failed to shutdown server: %v", err)
	}
}
