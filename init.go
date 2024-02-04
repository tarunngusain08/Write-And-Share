package main

import (
	"Write-And-Share/handlers"
	"Write-And-Share/repo"
	"Write-And-Share/service"
	"database/sql"
	"fmt"
)

const (
	userName = "tgusain"
	dbName   = "tgusain"
	disable  = "disable"
)

var handler *Handler

type Handler struct {
	*handlers.SignUpHandler
	*handlers.LoginHandler
	*handlers.GetNotesHandler
}

func initDB() (*sql.DB, error) {
	var err error
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", userName, dbName, disable)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")
	return db, nil
}

func init() {
	DB, err := initDB()
	if err != nil {
		panic(err)
	}
	signUpRepo := repo.NewSignUpRepo(DB)
	signUpHandler := handlers.NewSignUpHandler(signUpRepo)

	loginRepo := repo.NewLoginRepo(DB)
	loginService := service.NewLoginService(loginRepo)
	loginHandler := handlers.NewLoginHandler(loginService)

	getNotesRepo := repo.NewGetNotesRepo(DB)
	getNotesService := service.NewGetNotesRepo(getNotesRepo)
	getNotesHandler := handlers.NewGetNotesHandler(getNotesService)

	handler = new(Handler)
	handler.SignUpHandler = signUpHandler
	handler.LoginHandler = loginHandler
	handler.GetNotesHandler = getNotesHandler
}
