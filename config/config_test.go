package config

import "testing"

func TestDefaultConfig(t *testing.T) {
	conf := NewConfig(Config{})
	if conf.PixelsX != 500 && conf.PixelsY != 500 {
		t.Error("Error: PixelsX or PixelsY is not 500")
	}
	if conf.Ly != -250 {
		t.Error("Error LY is not properly defined")
	}
}

func TestOverrideConfig(t *testing.T) {
	conf := NewConfig(Config{})
	defaultD := conf.D
	defaultPixelsX := conf.PixelsX

	overrideD := defaultD + 1
	overridePixelsX := defaultPixelsX + 1
	conf = NewConfig(Config{
		D:       overrideD,
		PixelsX: overridePixelsX,
	})
	if conf.D != overrideD || conf.PixelsX != overridePixelsX || conf.Lx != -overridePixelsX/2 {
		t.Error("Error: overriding the config didn't work")
	}
}
