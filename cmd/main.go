package main

import (
	filmlib "github/film-lib"
	handler "github/film-lib/pkg/handler"
	"github/film-lib/pkg/repository"
	"github/film-lib/pkg/service"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s\n", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s\n", err.Error())
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
		log.Fatalf("error initializing db: %s\n", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	mux := handler.NewHandler(services)
	srv := new(filmlib.Server)

	if err := srv.Run(viper.GetString("port"), mux.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s\n", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
