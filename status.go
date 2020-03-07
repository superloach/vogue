package vogue

import "github.com/gdamore/tcell"

type Status struct {
	Screen tcell.Screen
	Style  tcell.Style
	Left   string
	Center string
	Right  string
}

func MkStatus(scr tcell.Screen, style tcell.Style) (*Status, error) {
	s := Status{}

	s.Screen = scr
	s.Style = style

	return &s, nil
}

func (s *Status) Fresh() {
	w, _ := s.Screen.Size()

	for i := 0; i < w; i++ {
		s.Screen.SetContent(i, 0, ' ', nil, s.Style)
	}

	for i, r := range []rune(s.Left) {
		s.Screen.SetContent(i, 0, r, nil, s.Style)
	}

	for i, r := range []rune(s.Center) {
		s.Screen.SetContent(
			(w-len(s.Center))/2+i,
			0, r, nil, s.Style,
		)
	}

	for i, r := range []rune(s.Right) {
		s.Screen.SetContent(
			w-len(s.Right)+i,
			0, r, nil, s.Style,
		)
	}
}
