package model

type Service[T any] interface{
	GetAll() []T
	GetById(id string) (T, error)
}