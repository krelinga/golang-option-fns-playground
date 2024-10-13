package main

import "testing"

func TestCWithSettings(t *testing.T) {
	t.Run("DirectSetA", func(t *testing.T) {
		c := &CImpl{}
		DirectSetA(c, 1)
		if c.a != 1 {
			t.Errorf("expected %d, got %d", 1, c.a)
		}
	})
	t.Run("CWithAnySettings", func(t *testing.T) {
		c := CWithAnySettings(SetA(1), SetB(2), SetC(3))
		if c.a != 1 {
			t.Errorf("expected %d, got %d", 1, c.a)
		}
		if c.b != 2 {
			t.Errorf("expected %d, got %d", 2, c.b)
		}
		if c.c != 3 {
			t.Errorf("expected %d, got %d", 3, c.c)
		}
	})
	t.Run("CWithAnySettingsUnkonwnSetting", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected panic")
			}
		}()
		CWithAnySettings("this is not a function")
	})
	t.Run("CWithTypedSettings", func(t *testing.T) {
		c := CWithTypedSettings(SetA(1), SetB(2), SetC(3))
		if c.a != 1 {
			t.Errorf("expected %d, got %d", 1, c.a)
		}
		if c.b != 2 {
			t.Errorf("expected %d, got %d", 2, c.b)
		}
		if c.c != 3 {
			t.Errorf("expected %d, got %d", 3, c.c)
		}
	})
	// t.Run("BWithTypedSettingsAndCOptionShouldNotCompile", func(t *testing.T) {
	// 	_ = BWithTypedSettings(SetA(1), SetB(2), SetC(3))
	// })
}
