package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type GenericGormRepository[T any] struct {
	DB *gorm.DB
}

func (repo *GenericGormRepository[T]) FindById(ctx context.Context, id int, target *T) error {
	tx := repo.DB.First(target, id)
	return tx.Error
}

func (repo *GenericGormRepository[T]) Save(ctx context.Context, target *T) error {
	result := repo.DB.Save(target)
	return result.Error

}

func (repo *GenericGormRepository[T]) FindByParams(ctx context.Context, target *[]T, queryParams map[string]interface{}) error {
	query := repo.DB.Model(target)
	for paramName, paramValue := range queryParams {
		query = query.Where(fmt.Sprintf("%v=?", paramName), paramValue)
	}
	res := repo.DB.Model(target).Where(query).Find(target)
	return res.Error
}

