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

type AApply interface {
	ApplyA(A)
	ApplyB(B)
	ApplyC(C)
}

type AOption func(A)

func (f AOption) ApplyA(a A) {
	f(a)
}

func (f AOption) ApplyB(b B) {
	f(b)
}

func (f AOption) ApplyC(c C) {
	f(c)
}

func SetA(in int) AOption {
	return func(a A) {
		a.SetA(in)
	}
}

func DirectSetA(a A, in int) {
	a.SetA(in)
}