package handlers

import (
	"Write-And-Share/contracts"
	"Write-And-Share/repo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	repo *repo.LoginRepo
}

func NewLoginHandler(loginRepo *repo.LoginRepo) *LoginHandler {
	return &LoginHandler{
		repo: loginRepo,
	}
}

func (l *LoginHandler) Login(ctx *gin.Context) {
	var details *contracts.UserDetails
	body, err := ctx.Request.GetBody()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = json.NewDecoder(body).Decode(&details)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = l.repo.Login(details)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
