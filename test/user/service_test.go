package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"

	"smanager/internal/user"
	"smanager/test/fixture"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	dbFixture := fixture.RunDBFixture()
	repo := dbFixture.UserRepo
	userService := user.NewUserService(repo)
	_, err := userService.CreateUser(context.TODO(), user.CreateUserDTO{"username", "password"})
	fmt.Println(err)
	assert.Nil(t, err)
	userList := make([]user.User, 0)
	paramMap := make(map[string]interface{})
	paramMap["username"] = "username"
	repo.FindByParams(context.TODO(), &userList, paramMap)
	assert.NotEmpty(t, userList, "Creating user")
	assert.Equal(t, "username", userList[0].Username, "Creating user")
	hasher := sha256.New()
	hasher.Write([]byte("password"))
	hashedPassword := (hasher.Sum(nil)[:])
	fmt.Println(hex.EncodeToString(hashedPassword))
	assert.Equal(t, hex.EncodeToString(hashedPassword), userList[0].Password)
}

func TestShouldnotCreateUserIfUsernameInUse(t *testing.T) {
	dbFixture := fixture.RunDBFixture()
	repo := dbFixture.UserRepo
	userService := user.NewUserService(repo)
	_, err := userService.CreateUser(context.TODO(), user.CreateUserDTO{fixture.TEST_USERNAME, "pass"})
	assert.NotNil(t, err)

}

func TestFindUser(t *testing.T) {
	dbFixture := fixture.RunDBFixture()
	repo := dbFixture.UserRepo
	userService := user.NewUserService(repo)
	storedUser, _ := userService.FindUserByUsername(context.TODO(), fixture.TEST_USERNAME)
	assert.Equal(t, fixture.TEST_USERNAME, storedUser.Username)
}
