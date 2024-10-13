package main

func CWithAnySettings(settings ...any) *CImpl {
	c := &CImpl{}
	for _, s := range settings {
		switch s := s.(type) {
		case AOptionFunc:
			s(c)
		case BOptionFunc:
			s(c)
		case COptionFunc:
			s(c)
		default:
			panic("unknown setting")
		}
	}
	return c
}

func CWithTypedSettings(settings ...COption) *CImpl {
	c := &CImpl{}
	for _, s := range settings {
		s.ApplyC(c)
	}
	return c
}

func BWithTypedSettings(settings ...BOption) *BImpl {
	b := &BImpl{}
	for _, s := range settings {
		s.ApplyB(b)
	}
	return b
}
