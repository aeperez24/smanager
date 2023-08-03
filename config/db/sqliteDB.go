package db

import (
	"smanager/internal/managedsecret"
	"smanager/internal/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbSqliteConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func DbSqliteConnectionWithFile(fileName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func Migrate(db *gorm.DB) {
	models := [...]interface{}{
		&user.User{},
		&managedsecret.ManagedSecret{},
	}
	for _, model := range models {
		db.AutoMigrate(model)
	}
}
