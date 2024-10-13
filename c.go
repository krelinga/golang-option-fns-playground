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

func SetC(in int) func(C) {
	return func(c C) {
		c.SetC(in)
	}
}
