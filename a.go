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

type AOption func(A)

func SetA(in int) AOption {
	return func(a A) {
		a.SetA(in)
	}
}

func DirectSetA(a A, in int) {
	a.SetA(in)
}