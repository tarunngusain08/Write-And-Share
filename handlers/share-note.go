package handlers

import (
	"Write-And-Share/contracts"
	"Write-And-Share/middlewares"
	"Write-And-Share/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShareNoteHandler struct {
	service *service.ShareNoteService
}

func NewShareNoteHandler(notesService *service.ShareNoteService) *ShareNoteHandler {
	return &ShareNoteHandler{
		notesService,
	}
}

func (g *ShareNoteHandler) ShareNote(ctx *gin.Context) {
	var note *contracts.ShareNoteRequest
	claims, err := middlewares.ExtractTokenClaimsFromHeader(ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	username := claims["username"].(string)
	if username == "" {
		ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	note.Username = username

	err = g.service.ShareNote(note)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, "Success")
}
