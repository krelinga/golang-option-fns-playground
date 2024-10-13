package main

type C interface {
	B
	SetC(c int)
}

type CImpl struct {
	BImpl
	c int
}

func (i *CImpl) SetC(c int) {
	i.c = c
}

type CApply interface {
	ApplyC(C)
}

type COption func(C)

func (f COption) ApplyC(c C) {
	f(c)
}

func SetC(in int) COption {
	return func(c C) {
		c.SetC(in)
	}
}
