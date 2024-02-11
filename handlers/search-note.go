package handlers

import (
	"Write-And-Share/contracts"
	"Write-And-Share/middlewares"
	"Write-And-Share/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SearchNoteHandler struct {
	service         *service.SearchNoteService
	defaultPageSize *int
}

func NewSearchNoteHandler(notesService *service.SearchNoteService, defaultPageSize *int) *SearchNoteHandler {
	return &SearchNoteHandler{
		notesService,
		defaultPageSize,
	}
}

func (g *SearchNoteHandler) Search(ctx *gin.Context) {
	var request *contracts.SearchNoteRequest
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
	if request.PageSize == nil {
		request.PageSize = g.defaultPageSize
	}

	noteList, err := g.service.SearchNote(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, noteList)
}
