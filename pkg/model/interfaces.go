package model

type IDGetter[T any] interface {
	GetById(id string) (T, error)
}

type Creator[T any] interface {
	Create(T) (error)
}
