package main

func CWithAnySettings(settings ...any) *CImpl {
	c := &CImpl{}
	for _, s := range settings {
		switch s := s.(type) {
		case AOption:
			s(c)
		case BOption:
			s(c)
		case COption:
			s(c)
		default:
			panic("unknown setting")
		}
	}
	return c
}

func CWithTypedSettings(settings ...CApply) *CImpl {
	c := &CImpl{}
	for _, s := range settings {
		s.ApplyC(c)
	}
	return c
}

func BWithTypedSettings(settings ...BApply) *BImpl {
	b := &BImpl{}
	for _, s := range settings {
		s.ApplyB(b)
	}
	return b
}