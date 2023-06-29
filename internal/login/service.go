package login

import (
	"context"
	"smanager/internal/user"
)

type LoginService struct {
	createTokenService CreateTokenService
	userService        UserService
}

type UserService interface {
	ValidateUsernameAndPassword(ctx context.Context, username, Password string) bool
	FindUserByUsername(ctx context.Context, username string) (*user.UserDTO, error)
}

type CreateTokenService interface {
	CreateToken(claims map[string]interface{}) (string, error)
}

func NewLoginService(userService UserService, createTokenService CreateTokenService) *LoginService {
	return &LoginService{userService: userService, createTokenService: createTokenService}
}

func (ls *LoginService) LoginWithUsernameAndPassword(ctx context.Context, username, password string) (string, error) {
	usernamePaswordMatch := ls.userService.ValidateUsernameAndPassword(ctx, username, password)
	if !usernamePaswordMatch {
		return "", newInvalidUserError()
	}
	userdto, er1 := ls.userService.FindUserByUsername(ctx, username)

	if er1 != nil {
		return "", er1
	}
	claimMap := make(map[string]interface{})
	claimMap["user"] = *userdto
	token, er1 := ls.createTokenService.CreateToken(claimMap)
	if er1 != nil {
		return "", er1
	}
	return token, nil
}
