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
	ApplyA(A)
	ApplyB(B)
	ApplyC(C)
}

type AOptionFunc func(A)

func (f AOptionFunc) ApplyA(a A) {
	f(a)
}

func (f AOptionFunc) ApplyB(b B) {
	f(b)
}

func (f AOptionFunc) ApplyC(c C) {
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
