package handlers

import (
	"Write-And-Share/contracts"
	"Write-And-Share/middlewares"
	"Write-And-Share/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UpsertNoteHandler struct {
	service *service.UpsertNoteService
}

func NewCreateNoteHandler(notesService *service.UpsertNoteService) *UpsertNoteHandler {
	return &UpsertNoteHandler{
		notesService,
	}
}

func (g *UpsertNoteHandler) UpsertNote(ctx *gin.Context) {
	var note *contracts.UpsertNoteRequest
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
	note.Time = time.Now()

	err = g.service.UpsertNote(note)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, "Success")
}
