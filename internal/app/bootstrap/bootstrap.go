package bootstrap

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"online-shop-backend/internal/app/transport/http/handler"
	"online-shop-backend/internal/config"
	"online-shop-backend/internal/repository"
	"online-shop-backend/internal/service"
	"online-shop-backend/pkg/db"
	"online-shop-backend/pkg/logger"
	"online-shop-backend/pkg/server"
	"os"
	"os/signal"
	"syscall"
)

func Website() {
	logger.InitLogger()

	cnf, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	validate := validator.New()

	db, err := db.NewPostgresDB(cnf)
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("error with connecting to database: %s", err.Error()))
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services, validate)
	srv := new(server.Server)

	go func() {
		if err := srv.Run(cnf.HTTPPort, handlers.InitRoutes()); err != nil {
			zap.L().Fatal(fmt.Sprintf("error with running server: %s", err.Error()))
		}
	}()

	zap.L().Info("Online Shop Backend started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	zap.L().Info("Online Shop Backend Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		zap.L().Fatal(fmt.Sprintf("error with shutting down server: %s", err.Error()))
	}

	if err := db.Close(); err != nil {
		zap.L().Fatal(fmt.Sprintf("error with closing db: %s", err.Error()))
	}
}
