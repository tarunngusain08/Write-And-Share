package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func main() {
	r := gin.Default()

	limiter := ratelimit.NewBucketWithRate(100, 100)
	limiter.Capacity()

	r.Use(func(c *gin.Context) {
		limiter.Wait(1)
		c.Next()
	})

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/signup", handler.SignUpHandler.SignUp)
		auth.POST("/login", handler.LoginHandler.Login)
	}

	notes := api.Group("/notes")
	{
		notes.POST("")
		notes.GET("/:id", handler.GetNotesHandler.GetNotesById)
		notes.GET("", handler.GetNotesHandler.GetAllNotes)
		notes.POST("", handler.UpsertNoteHandler.UpsertNote)
		notes.PUT("/:id", handler.UpsertNoteHandler.UpsertNote)
		notes.DELETE("/:id", handler.DeleteNoteHandler.DeleteNote)
		notes.POST("/:id/share", handler.ShareNoteHandler.ShareNote)
		notes.GET("/search?q=:query", handler.SearchNoteHandler.Search)
	}
}
