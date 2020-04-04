package vogue

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
)

type Vogue struct {
	Screen tcell.Screen

	Config Config

	Status *Status
	Tabs   *Tabs
}

func NewVogue() (*Vogue, error) {
	v := &Vogue{}

	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	v.Screen = screen
	v.Screen.EnableMouse()

	cfg, err := LoadConfig()
	if err != nil {
		cfg = make(Config)
	}
	v.Config = cfg

	v.Status = &Status{
		Left: func() string {
			buf := v.Tabs.Buffer()
			return fmt.Sprintf("%d:%d", buf.Line, buf.Col)
		},
		Center: func() string {
			return "vogue"
		},
		Right: func() string {
			return time.Now().Format(time.Kitchen)
		},
	}

	v.Tabs = NewTabs()

	return v, nil
}
