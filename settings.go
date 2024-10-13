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
