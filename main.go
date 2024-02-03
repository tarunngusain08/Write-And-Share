package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	api := r.Group("/api")

	auth := api.Group("/auth")
	auth.POST("/signup", handler.SignUpHandler.SignUp)
	auth.POST("/login", handler.LoginHandler.Login)

	notes := api.Group("/notes")
	notes.POST("")
	notes.GET("/:id", handler.GetNotesHandler.GetNotesById)
	notes.GET("", handler.GetNotesHandler.GetAllNotes)
}
