package app

import (
	"context"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"shop/internal/config"
	controller "shop/internal/controller/v1"
	"shop/internal/driver"
	"shop/internal/repository"
	"shop/internal/server"
	"shop/internal/service"
	"syscall"
	"time"
)

func Run(cfg config.Config) error {
	db, err := driver.NewPostgresPool(cfg.PSQL)
	if err != nil {
		return err
	}
	defer func() {
		db.Close()
		log.Info().Msg("Connection with db close")
	}()

	repo := repository.NewStockPostgres(db)
	services := service.NewStock(repo)
	handler := controller.New(services)

	srv := new(server.Server)
	defer func() {
		log.Info().Msg("App Shutting Down")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = srv.Shutdown(ctx)
		if err != nil {
			log.Error().Err(err)
		}
		log.Info().Msg("Server Stopped")
	}()

	errChan := make(chan error, 1)

	go func() {
		if err = srv.Run(cfg.Server.Port, handler.InitRouter()); err != nil {
			errChan <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	select {
	case <-quit:
	case err = <-errChan:
		return err
	}

	return nil
}
