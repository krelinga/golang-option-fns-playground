package main

func CWithAnySettings(settings ...any) *CImpl {
	c := &CImpl{}
	for _, s := range settings {
		switch s := s.(type) {
		case func(A):
			s(c)
		case func(B):
			s(c)
		case func(C):
			s(c)
		default:
			panic("unknown setting")
		}
	}
	return c
}
