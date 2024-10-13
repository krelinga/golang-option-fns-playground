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

type BApply interface {
	ApplyB(B)
}

type BOption func(B)

func (f BOption) ApplyB(b B) {
	f(b)
}

func (f BOption) ApplyC(c C) {
	f(c)
}

func SetB(in int) BOption {
	return func(b B) {
		b.SetB(in)
	}
}
