package vogue

import "github.com/gdamore/tcell"

func (v *Vogue) Fresh() error {
	style := v.Config.GetStyle(
		"status",
		tcell.StyleDefault.Background(tcell.ColorWhite),
	)
	err := v.Status.Fresh(v.Screen, style)
	if err != nil {
		return err
	}

	v.Screen.Show()

	return nil
}
