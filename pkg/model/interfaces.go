package model

type IDGetter[T any] interface {
	GetById(id string) (T, error)
}

type Creator[Input, T any] interface {
	Create(input Input) (T, error)
}

type IDGetterCreator[Input, T any] interface {
	IDGetter[T]
	Creator[Input, T]
}