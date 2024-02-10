package handlers

import (
	"Write-And-Share/contracts"
	"Write-And-Share/middlewares"
	"Write-And-Share/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeleteNoteHandler struct {
	service *service.DeleteNoteService
}

func NewDeleteNoteHandler(notesService *service.DeleteNoteService) *DeleteNoteHandler {
	return &DeleteNoteHandler{
		notesService,
	}
}

func (g *DeleteNoteHandler) DeleteNote(ctx *gin.Context) {
	var note *contracts.DeleteNoteRequest
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

	err = g.service.DeleteNote(note)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, "Success")
}
