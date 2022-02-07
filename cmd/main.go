package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/abbos-ron2/go_postgresql/api"
	"github.com/abbos-ron2/go_postgresql/config"
	"github.com/abbos-ron2/go_postgresql/service"
	"github.com/abbos-ron2/go_postgresql/storage"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	cfg := config.Load()

	psqlConnString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	// db, err := sqlx.Connect("postgres", psqlConnString)
	db, err := sql.Open("postgres", psqlConnString)
	if err != nil {
		log.Panic("error connecting to postgres", err)
	}

	storage := storage.NewStorage(db)

	serviceManager := service.NewServiceManager(storage)

	router := api.New(serviceManager)

	router.Run()
}
