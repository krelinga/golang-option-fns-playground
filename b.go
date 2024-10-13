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

func SetB(in int) func(B) {
	return func(b B) {
		b.SetB(in)
	}
}
