package handlers

import (
	"Write-And-Share/contracts"
	"Write-And-Share/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpHandler struct {
	signupService *service.SignupService
}

func NewSignUpHandler(signupService *service.SignupService) *SignUpHandler {
	return &SignUpHandler{
		signupService: signupService,
	}
}

func (s *SignUpHandler) SignUp(ctx *gin.Context) {

	var details *contracts.LoginRequest
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

	token, err := s.signupService.Signup(details)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}
