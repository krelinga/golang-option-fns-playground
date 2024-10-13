package main

type A interface {
	SetA(a int)
}

type AImpl struct {
	a int
}

func (i *AImpl) SetA(a int) {
	i.a = a
}

type AOption interface {
	ApplyToA(A)
	BOption
}

type AOptionFunc func(A)

func (f AOptionFunc) ApplyToA(a A) {
	f(a)
}

func (f AOptionFunc) ApplyToB(b B) {
	f(b)
}

func (f AOptionFunc) ApplyToC(c C) {
	f(c)
}

func SetA(in int) AOptionFunc {
	return func(a A) {
		a.SetA(in)
	}
}

func DirectSetA(a A, in int) {
	a.SetA(in)
}
