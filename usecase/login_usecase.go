package usecase

import (
	"gin-clean-arch/bootstrap"
	"gin-clean-arch/domain"
	tokenutil "gin-clean-arch/internal"
)

type LoginUsecase struct {
	userRepository domain.UserRepository
	env            *bootstrap.Env
}

func NewLoginUsecase(userRepository domain.UserRepository, env *bootstrap.Env) domain.LoginUsecase {
	return &LoginUsecase{
		userRepository: userRepository,
		env:            env,
	}
}

func (lu *LoginUsecase) GetUserByEmail(email string) (*domain.User, error) {
	return lu.userRepository.FindByEmail(email)
}

func (lu *LoginUsecase) CreateAccessToken(user *domain.User) (string, error) {
	return tokenutil.CreateAccessToken(user, lu.env.SecretKey, lu.env.AccessTokenExpiryHour)
}

func (lu *LoginUsecase) CreateRefreshToken(user *domain.User) (string, error) {
	return tokenutil.CreateRefreshToken(user, lu.env.SecretKey, lu.env.AccessTokenExpiryHour)
}
