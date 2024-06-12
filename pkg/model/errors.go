package model

type ErrNotFound struct {
	Err error
}

func (n ErrNotFound) Error() string {
	return n.Err.Error()
}
