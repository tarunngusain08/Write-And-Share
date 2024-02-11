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

var (
	handler         *Handler
	defaultPageSize = 10
)

type Handler struct {
	*handlers.SignUpHandler
	*handlers.LoginHandler
	*handlers.GetNotesHandler
	*handlers.UpsertNoteHandler
	*handlers.DeleteNoteHandler
	*handlers.ShareNoteHandler
	*handlers.SearchNoteHandler
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
	loginRepo := repo.NewLoginRepo(DB)
	getNotesRepo := repo.NewGetNotesRepo(DB)
	upsertNoteRepo := repo.NewUpsertNoteRepo(DB)
	deleteNoteRepo := repo.NewDeleteNoteRepo(DB)
	shareNoteRepo := repo.NewShareNoteRepo(DB)
	searchNoteRepo := repo.NewSearchNoteRepo(DB)

	loginService := service.NewLoginService(loginRepo)
	signUpService := service.NewSignupService(signUpRepo, loginService)
	getNotesService := service.NewGetNotesRepo(getNotesRepo)
	upsertNoteService := service.NewUpsertNoteService(upsertNoteRepo)
	deleteNoteService := service.NewDeleteNoteService(deleteNoteRepo)
	shareNoteService := service.NewShareNoteService(shareNoteRepo)
	searchNoteService := service.NewSearchNoteService(searchNoteRepo)

	signUpHandler := handlers.NewSignUpHandler(signUpService)
	loginHandler := handlers.NewLoginHandler(loginService)
	getNotesHandler := handlers.NewGetNotesHandler(getNotesService, &defaultPageSize)
	upsertNoteHandler := handlers.NewUpsertNoteHandler(upsertNoteService)
	deleteNoteHandler := handlers.NewDeleteNoteHandler(deleteNoteService)
	shareNoteHandler := handlers.NewShareNoteHandler(shareNoteService)
	searchNoteHandler := handlers.NewSearchNoteHandler(searchNoteService, &defaultPageSize)

	handler = new(Handler)
	handler.SignUpHandler = signUpHandler
	handler.LoginHandler = loginHandler
	handler.GetNotesHandler = getNotesHandler
	handler.UpsertNoteHandler = upsertNoteHandler
	handler.DeleteNoteHandler = deleteNoteHandler
	handler.ShareNoteHandler = shareNoteHandler
	handler.SearchNoteHandler = searchNoteHandler
}
