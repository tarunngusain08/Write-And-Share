package handlers

import (
	"Write-And-Share/contracts"
	"Write-And-Share/middlewares"
	"Write-And-Share/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetNotesHandler struct {
	service *service.GetNotesService
}

func NewGetNotesHandler(notesService *service.GetNotesService) *GetNotesHandler {
	return &GetNotesHandler{
		notesService,
	}
}

func (g *GetNotesHandler) GetAllNotes(ctx *gin.Context) {
	var user *contracts.GetNotesRequest
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
	user.UserName = username

	notes, err := g.service.GetNotes(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, notes)
}

func (g *GetNotesHandler) GetNotesById(ctx *gin.Context) {
	var user *contracts.GetNotesRequest
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

	user.UserName = username
	noteId := ctx.Param("id")
	user.NoteId = &noteId

	notes, err := g.service.GetNotes(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, notes)
}
