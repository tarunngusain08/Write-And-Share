package handlers

import (
	"Write-And-Share/contracts"
	"Write-And-Share/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	*service.LoginService
}

func NewLoginHandler(loginService *service.LoginService) *LoginHandler {
	return &LoginHandler{
		loginService,
	}
}

func (l *LoginHandler) Login(ctx *gin.Context) {
	var user *contracts.LoginRequest
	body, err := ctx.Request.GetBody()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	err = json.NewDecoder(body).Decode(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	token, err := l.LoginService.Login(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}
