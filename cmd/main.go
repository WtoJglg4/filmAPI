package main

import (
	"context"
	filmapi "github/film-lib"
	handler "github/film-lib/pkg/handler"
	"github/film-lib/pkg/repository"
	"github/film-lib/pkg/service"
	"os"
	"os/signal"
	"syscall"

	_ "github/film-lib/docs"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//	@title			FilmAPI
//	@version		1.0
//	@description	backend of the application, which provides a REST API for managing the films database

//	@host		localhost:3000
//	@BasePath	/

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description					Use "Bearer " + token. Admin: {login: "admin", pass: "admin"}
func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s\n", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s\n", err.Error())
	}

	dbConfig := repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		logrus.Fatalf("error initializing db: %s\n", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	mux := handler.NewHandler(services)
	srv := new(filmapi.Server)

	services.Authorization.CreateUser(filmapi.User{
		Username: "admin",
		Password: "admin",
		Role:     "admin",
	})

	go func() {
		if err := srv.Run(viper.GetString("port"), mux.InitRoutes()); err != nil {
			logrus.Fatalf("error while running http server: %s\n", err.Error())
		}
	}()

	logrus.Printf("FilmAPI started on :%s\n", viper.GetString("port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Println("FilmAPI Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
