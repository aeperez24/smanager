package repository

import "context"

type GenericRepository[T any] interface {
	FindById(ctx context.Context, id int, target *T) error
	Save(ctx context.Context, target *T) error
	FindByParams(ctx context.Context, target *[]T, queryParams map[string]interface{}) error
}

type RepoQuery struct {
	Key   string
	Value interface{}
}

type repoQueryBuilder struct {
	queryMap map[string]interface{}
}

func QueriBuilder() *repoQueryBuilder {
	qmap := make(map[string]interface{})
	return &repoQueryBuilder{
		queryMap: qmap,
	}
}
func (builder *repoQueryBuilder) With(key string, value interface{}) *repoQueryBuilder {
	builder.queryMap[key] = value

	return builder
}

func (builder *repoQueryBuilder) Build() map[string]interface{} {
	return builder.queryMap
}
