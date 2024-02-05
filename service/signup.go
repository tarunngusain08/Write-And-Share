package service

import (
	"Write-And-Share/contracts"
	"Write-And-Share/repo"
)

type SignupService struct {
	signUpRepo   *repo.SignUpRepo
	loginService *LoginService
}

func NewSignupService(signupRepo *repo.SignUpRepo, service *LoginService) *SignupService {
	return &SignupService{
		signupRepo,
		service,
	}
}

func (s *SignupService) Signup(user *contracts.LoginRequest) (string, error) {
	err := s.signUpRepo.SignUp(user)
	if err != nil {
		return "", err
	}
	return s.loginService.Login(user)
}
