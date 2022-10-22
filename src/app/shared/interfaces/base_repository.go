package interfaces

import (
	"base-fiber-api/src/app/shared/pkg"
)

type BaseRepository[T interface{}] interface {
	List(pagination pkg.Pagination) (*pkg.Pagination, error)
	Get(id string) (*T, error)
	Store(model *T) (*T, error)
	Edit(model *T) (*T, error)
	Delete(model *T) error
	FindBy(field string, value string) (*T, error)
	FindManyBy(field []string, value string) (*T, error)
}
