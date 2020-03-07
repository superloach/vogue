package vogue

import "github.com/gdamore/tcell"

func (v *Vogue) Fresh() {
	v.Screen.Clear()
	v.Status.Fresh()
}

func (v *Vogue) Run() error {
	for {
		v.Screen.Show()

		ev := v.Screen.PollEvent()
		if ev == nil {
			return nil
		}

		switch ev.(type) {
		case *tcell.EventError:
			return ev.(error)
		case *tcell.EventInterrupt:
			v.Fresh()
		case *tcell.EventKey:
			evk := ev.(*tcell.EventKey)

			switch evk.Key() {
			case tcell.KeyESC:
				v.Fini()
			default:
			}
		case *tcell.EventResize:
			v.Fresh()
		default:
		}
	}

	return nil
}
