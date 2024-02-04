package service

import (
	"Write-And-Share/contracts"
	"Write-And-Share/repo"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type LoginService struct {
	*repo.LoginRepo
}

func NewLoginService(loginRepo *repo.LoginRepo) *LoginService {
	return &LoginService{
		loginRepo,
	}
}

func (l *LoginService) Login(user *contracts.LoginRequest) (string, error) {
	err := l.LoginRepo.Login(user)
	if err != nil {
		return "", err
	}
	return l.tokenGenerator(user.UserName)
}

func (l *LoginService) tokenGenerator(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
