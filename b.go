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
	ApplyToB(B)
	COption
}

type BOptionFunc func(B)

func (f BOptionFunc) ApplyToB(b B) {
	f(b)
}

func (f BOptionFunc) ApplyToC(c C) {
	f(c)
}

func SetB(in int) BOptionFunc {
	return func(b B) {
		b.SetB(in)
	}
}
