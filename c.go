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

type COption interface {
	ApplyToC(C)
}

type COptionFunc func(C)

func (f COptionFunc) ApplyToC(c C) {
	f(c)
}

func SetC(in int) COptionFunc {
	return func(c C) {
		c.SetC(in)
	}
}
