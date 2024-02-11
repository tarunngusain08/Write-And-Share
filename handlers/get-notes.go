package handlers

import (
	"Write-And-Share/contracts"
	"Write-And-Share/middlewares"
	"Write-And-Share/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetNotesHandler struct {
	service         *service.GetNotesService
	defaultPageSize *int
}

func NewGetNotesHandler(notesService *service.GetNotesService, defaultPageSize *int) *GetNotesHandler {
	return &GetNotesHandler{
		notesService,
		defaultPageSize,
	}
}

func (g *GetNotesHandler) GetAllNotes(ctx *gin.Context) {
	var request *contracts.GetNotesRequest
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
	request.UserName = username
	if request.PageSize == nil {
		request.PageSize = g.defaultPageSize
	}

	notes, err := g.service.GetNotes(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, notes)
}

func (g *GetNotesHandler) GetNotesById(ctx *gin.Context) {
	var request *contracts.GetNotesRequest
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

	request.UserName = username
	noteId := ctx.Param("id")
	request.NoteId = &noteId

	notes, err := g.service.GetNotes(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, notes)
}
