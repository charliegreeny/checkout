package model

type Service[T any] interface{
	GetById(id string) (T, error)
}