package repository

type Adapter interface {
}

type adapter struct {
}

func New() Adapter {
    return &adapter{}
}
