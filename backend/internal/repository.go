package internal

type IRepository[T any] interface {
	Create(schema T) error
	Read() error
	Update(schema T) error
	Delete(id int) error
}
