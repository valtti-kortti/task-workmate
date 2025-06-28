package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"task-workmate/internal/api"
	"task-workmate/internal/config"
	"task-workmate/internal/repository"
	"task-workmate/internal/service"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	if err := godotenv.Load(config.EnvPath); err != nil {
		log.Fatal(err)
		return
	}

	var cfg config.AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
		return
	}

	rep := repository.NewRepository()

	serv := service.NewService(rep)

	app := api.NewRouters(&api.Routers{Service: serv})

	go func() {
		if err := app.Listen(cfg.Rest.ListenAddress); err != nil {
			log.Fatal(errors.New(err.Error()))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

}
