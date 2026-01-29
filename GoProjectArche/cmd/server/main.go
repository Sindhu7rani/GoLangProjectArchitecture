package main

import (
	"log"
	"net/http"

	"GoProjectArche/config"
	"GoProjectArche/internal/db"
	"GoProjectArche/internal/handler"
	"GoProjectArche/internal/repository"
	"GoProjectArche/internal/router"
	"GoProjectArche/internal/service"
)

func main() {

	//load application config
	cfg := config.LoadConfig()

	//connect db
	/*make DB connection after config because,
	for DB connection we want data from config*/
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	//initialize repo which communicate with DB
	repo := &repository.AccountRepository{DB: database}
	//initialize service which communicate with repo
	service := &service.AccountService{Repo: repo}
	//initialize handler which communicate with service
	handler := &handler.AccountHandler{Service: service}

	/* Router maps URL paths to handlers
	Middleware (authentication, logging, etc.) is also attached here*/
	//setup router
	r := router.SetupRoutes(handler)

	//start server
	log.Println("Server started on port", cfg.Port)
	// We pass `r` because it contains all routes and middleware
	err = http.ListenAndServe(":"+cfg.Port, r)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
