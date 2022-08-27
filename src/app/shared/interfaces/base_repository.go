package interfaces

type BaseRepository[T interface{}] interface {
	List() ([]T, error)
	Get(id string) (*T, error)
	Store(model *T) (*T, error)
	Edit(model *T) (*T, error)
	Delete(model *T) error
}
