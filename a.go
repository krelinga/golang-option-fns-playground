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

func SetA(in int) func(A) {
	return func(a A) {
		a.SetA(in)
	}
}

func DirectSetA(a A, in int) {
	a.SetA(in)
}