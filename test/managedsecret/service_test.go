package managedsecret

import (
	"context"
	"fmt"
	"smanager/internal/managedsecret"
	"smanager/internal/user"
	"smanager/test/fixture"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateManagedSecret(t *testing.T) {
	dbFixture := fixture.RunDBFixture()
	repo := dbFixture.ManagedSecretRepo
	mgService := &managedsecret.ManagedSecretService{ManagedSecretRepo: repo}
	ctx := context.WithValue(context.TODO(), "user", user.UserDTO{
		Id:       1,
		Username: fixture.TEST_USERNAME,
	})
	secretName := "the secret Name"
	secretValue := "the secret Value"
	mgService.CreateManagedSecret(ctx, secretName, secretValue)
	queryResult := make([]managedsecret.ManagedSecret, 0)
	queryParam := make(map[string]interface{})
	queryParam["name"] = secretName
	queryParam["user_id"] = 1
	repo.FindByParams(context.TODO(), &queryResult, queryParam)
	fmt.Println(queryResult[0])
	assert.Equal(t, secretName, queryResult[0].Name)
	assert.Equal(t, secretValue, queryResult[0].Value)
}

func TestListManagedSecrets(t *testing.T) {
	dbFixture := fixture.RunDBFixture()
	repo := dbFixture.ManagedSecretRepo
	mgService := &managedsecret.ManagedSecretService{ManagedSecretRepo: repo}
	ctx := context.WithValue(context.TODO(), "user", user.UserDTO{
		Id:       1,
		Username: fixture.TEST_USERNAME,
	})
	managedSecretList, err := mgService.ListManagedSecret(ctx)
	assert.Nil(t, err)
	assert.Contains(t, managedSecretList, managedsecret.ManagedSecretDto{
		ID:   1,
		Name: fixture.TEST_SECRET_NAME_1,
	})
	assert.Contains(t, managedSecretList, managedsecret.ManagedSecretDto{
		ID:   2,
		Name: fixture.TEST_SECRET_NAME_2,
	})
	assert.Len(t, managedSecretList, 2)

}

func TestGetManagedSecrets(t *testing.T) {
	dbFixture := fixture.RunDBFixture()
	repo := dbFixture.ManagedSecretRepo
	mgService := &managedsecret.ManagedSecretService{ManagedSecretRepo: repo}
	ctx := context.WithValue(context.TODO(), "user", user.UserDTO{
		Id:       1,
		Username: fixture.TEST_USERNAME,
	})
	secret, err := mgService.GetSecret(ctx, fixture.TEST_SECRET_NAME_1)
	assert.Nil(t, err)
	assert.Equal(t, fixture.TEST_SECRET_VALUE_1, secret)

}
