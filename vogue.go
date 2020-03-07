package vogue

import "github.com/gdamore/tcell"

type Vogue struct {
	Screen tcell.Screen
	Status *Status
}

func MkVogue() (*Vogue, error) {
	v := Vogue{}

	scr, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	v.Screen = scr

	status, err := MkStatus(
		v.Screen,
		tcell.StyleDefault.
			Foreground(tcell.ColorBlack).
			Background(tcell.ColorWhite),
	)
	if err != nil {
		return nil, err
	}
	v.Status = status

	return &v, nil
}

func (v *Vogue) Init() error {
	err := v.Screen.Init()
	if err != nil {
		return err
	}

	v.Screen.Clear()

	return nil
}

func (v *Vogue) Fini() {
	v.Screen.Fini()
}
