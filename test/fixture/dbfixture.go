package fixture

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"smanager/config/db"
	"smanager/config/repository"
	"smanager/internal/managedsecret"
	"smanager/internal/user"

	"gorm.io/gorm"
)

const TEST_USERNAME string = "usernameForTests"
const TEST_PASSWORD string = "passwordForTests"
const TEST_ID int = 1
const TEST_SECRET_NAME_1 string = "secretName1"
const TEST_SECRET_VALUE_1 string = "secretValue1"

const TEST_SECRET_NAME_2 string = "secretName2"
const TEST_SECRET_VALUE_2 string = "secretValue2"

func RunDBFixture() DBFixture {
	DB := db.DbSqliteConnection()
	prepareDB(DB)
	userRepo := &repository.GenericGormRepository[user.User]{DB: DB}
	ManagedSecretRepo := &repository.GenericGormRepository[managedsecret.ManagedSecret]{DB: DB}
	return DBFixture{UserRepo: userRepo, ManagedSecretRepo: ManagedSecretRepo}

}

func prepareDB(DB *gorm.DB) {
	db.Migrate(DB)
	prepareUsers(DB)
	prepareManagedSecrets(DB)
}

func prepareUsers(DB *gorm.DB) {
	hasher := sha256.New()
	hasher.Write([]byte(TEST_PASSWORD))
	userRepo := &repository.GenericGormRepository[user.User]{DB: DB}

	userRepo.Save(context.TODO(), &user.User{
		Username: TEST_USERNAME,
		Password: hex.EncodeToString(hasher.Sum(nil)[:]),
		Enabled:  true,
	})
}

func prepareManagedSecrets(DB *gorm.DB) {
	ManagedSecretRepo := &repository.GenericGormRepository[managedsecret.ManagedSecret]{DB: DB}
	ManagedSecretRepo.Save(context.TODO(), &managedsecret.ManagedSecret{
		UserId: uint(TEST_ID),
		Name:   TEST_SECRET_NAME_1,
		Value:  TEST_SECRET_VALUE_1,
	})
	ManagedSecretRepo.Save(context.TODO(), &managedsecret.ManagedSecret{
		UserId: uint(TEST_ID),
		Name:   TEST_SECRET_NAME_2,
		Value:  TEST_SECRET_VALUE_2,
	})
}

type DBFixture struct {
	UserRepo          repository.GenericRepository[user.User]
	ManagedSecretRepo repository.GenericRepository[managedsecret.ManagedSecret]
}
