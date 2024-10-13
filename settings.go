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
		s.ApplyToC(c)
	}
	return c
}

func BWithTypedSettings(settings ...BOption) *BImpl {
	b := &BImpl{}
	for _, s := range settings {
		s.ApplyToB(b)
	}
	return b
}
