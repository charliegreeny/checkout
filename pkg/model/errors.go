package model

type NotFoundErr struct {
	Err error
}

func (n NotFoundErr) Error() string {
	return n.Err.Error()
}
