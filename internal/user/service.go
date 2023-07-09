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
		return nil, fmt.Errorf("CreateUser:%w", err)
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
	user, err := u.findUserByUsername(ctx, username)
	if(err!=nil){
		return nil,fmt.Errorf("FindUserByUsername:%w", err)
	}
	if user == nil {
		return nil, nil
	}
	return &UserDTO{
		Id:       int(user.ID),
		Username: user.Username,
	}, err
}

func (u *UserService) ValidateUsernameAndPassword(ctx context.Context, username, password string) bool {
	user, err := u.findUserByUsername(ctx, username)
	if err != nil {
		return false
	}
	if user == nil {
		return false
	}
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := (hasher.Sum(nil)[:])
	return string(hashedPassword) == user.Password
}

func (u *UserService) IsValidUser(ctx context.Context, username string) (bool, error) {
	res, err := u.FindUserByUsername(ctx, username)
	return res != nil, fmt.Errorf("isValidUser:%w", err)
}

func (u *UserService) findUserByUsername(ctx context.Context, username string) (*User, error) {
	userList := make([]User, 0)
	builder := repository.QueriBuilder().
		With("username", username).
		With("enabled", true)

	err := u.userRepo.FindByParams(ctx, &userList, builder.Build())
	if err != nil {
		return nil, fmt.Errorf("findUserByUsername: %w", err)
	}

	if len(userList) == 0 {
		return nil, nil
	}

	return &userList[0], nil
}
