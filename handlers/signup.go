package handlers

import (
	"Write-And-Share/contracts"
	"Write-And-Share/repo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpHandler struct {
	repo *repo.SignUpRepo
}

func NewSignUpHandler(signUpRepo *repo.SignUpRepo) *SignUpHandler {
	return &SignUpHandler{
		repo: signUpRepo,
	}
}

func (s *SignUpHandler) SignUp(ctx *gin.Context) {

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

	err = s.repo.SignUp(details)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, "Success")
}
