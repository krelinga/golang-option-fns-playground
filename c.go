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

type COption func(C)

func SetC(in int) COption {
	return func(c C) {
		c.SetC(in)
	}
}
