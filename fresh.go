package vogue

import "github.com/gdamore/tcell"

func (v *Vogue) Fresh() {
	v.Status.Fresh(v.Screen, tcell.StyleDefault)
	v.Tabs.Fresh(v.Screen, tcell.StyleDefault, tcell.StyleDefault)
	v.Tabs.Buffer().Fresh(v.Screen)

	v.Screen.Show()
}
