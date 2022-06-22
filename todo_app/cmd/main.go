package main

import (
	"log"

	todoapp "github.com/ixursandbek/todo-app"
	"github.com/ixursandbek/todo-app/pkg/handler"
	"github.com/ixursandbek/todo-app/pkg/repository"
	"github.com/ixursandbek/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("xatolik config initsializatsiyasida: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todoapp.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatalf("http serverini ishga tushirishda xatolik yuz berdi: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
