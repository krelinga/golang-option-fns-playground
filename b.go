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

type BOption interface {
	ApplyB(B)
}

type BOptionFunc func(B)

func (f BOptionFunc) ApplyB(b B) {
	f(b)
}

func (f BOptionFunc) ApplyC(c C) {
	f(c)
}

func SetB(in int) BOptionFunc {
	return func(b B) {
		b.SetB(in)
	}
}
