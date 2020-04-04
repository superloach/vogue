package vogue

import "github.com/gdamore/tcell"

func (v *Vogue) Run() error {
	err := v.Screen.Init()
	if err != nil {
		return err
	}

	for {
		ev := v.Screen.PollEvent()

		switch ev.(type) {
		case *tcell.EventError:
			return ev.(error)
		case *tcell.EventInterrupt, *tcell.EventResize:
			v.Fresh()
		case *tcell.EventKey:
			go v.Key(ev.(*tcell.EventKey))
		case *tcell.EventMouse:
			go v.Mouse(ev.(*tcell.EventMouse))
		case eventQuit:
			return nil
		default:
		}
	}

	return nil
}

func (v *Vogue) Fini() {
	v.Screen.Fini()

	err := v.Config.Save()
	if err != nil {
		panic(err)
	}
}

func (v *Vogue) Key(evk *tcell.EventKey) {
	switch evk.Key() {
	case tcell.KeyCtrlQ:
		v.Screen.PostEventWait(eventQuit{})
	case tcell.KeyESC:
		v.Fresh()
		v.Screen.Sync()
	case tcell.KeyCtrlW, tcell.KeyCtrlJ, tcell.KeyCtrlK:
		v.Tabs.Key(evk)
	default:
		v.Tabs.Buffer().Key(evk)
	}
}

func (v *Vogue) Mouse(evm *tcell.EventMouse) {
	panic("vogue mouse stub")
}
