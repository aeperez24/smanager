package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"smanager/config/repository"
)

type UserService struct {
	userRepo repository.GenericRepository[User]
}

func NewUserService(repo repository.GenericRepository[User]) *UserService {
	return &UserService{userRepo: repo}
}

type CreateUserDTO struct {
	Username string
	Password string
}
type UserDTO struct {
	Id       int
	Username string
}

func (u *UserService) CreateUser(ctx context.Context, req CreateUserDTO) (*UserDTO, error) {
	hasher := sha256.New()
	hasher.Write([]byte(req.Password))
	hashedPassword := (hasher.Sum(nil)[:])

	alreadyStoredUser, err := u.FindUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if alreadyStoredUser != nil {
		return nil, fmt.Errorf("username %v already exists", req.Username)
	}
	userToSave := &User{
		Username: req.Username,
		Password: hex.EncodeToString(hashedPassword),
		Enabled:  true,
	}
	u.userRepo.Save(ctx, userToSave)
	return &UserDTO{
		Id:       int(userToSave.ID),
		Username: req.Username,
	}, nil
}

func (u *UserService) FindUserByUsername(ctx context.Context, username string) (*UserDTO, error) {
	userList := make([]User, 0)
	builder := repository.QueriBuilder().
		With("username", username).
		With("enabled", true)

	err := u.userRepo.FindByParams(ctx, &userList, builder.Build())
	if err != nil {
		return nil, err
	}

	if len(userList) == 0 {
		return nil, nil
	}

	return &UserDTO{
		Id:       int(userList[0].ID),
		Username: userList[0].Username,
	}, nil
}

func (u *UserService) ValidateUsernameAndPassword(ctx context.Context, username, Password string) bool {
	return false
}

func (u *UserService) IsValidUser(ctx context.Context, username string) (bool, error) {
	res, err := u.FindUserByUsername(ctx, username)
	return res != nil, err
}
