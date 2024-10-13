package main

type B interface {
	A
	SetB(b int)
}

type BImpl struct {
	AImpl
	b int
}

func (i *BImpl) SetB(b int) {
	i.b = b
}

type BOption func(B)

func SetB(in int) BOption {
	return func(b B) {
		b.SetB(in)
	}
}
