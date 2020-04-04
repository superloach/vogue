package vogue

import "github.com/gdamore/tcell"

type Status struct {
	Left   func() string
	Center func() string
	Right  func() string
}

func (s *Status) Fresh(screen tcell.Screen, style tcell.Style) error {
	var left, center, right string

	w, _ := screen.Size()

	if s.Left != nil {
		left = s.Left()
	}
	loff := 1

	if s.Center != nil {
		center = s.Center()
	}
	coff := (w / 2) - (len(center) / 2)

	if s.Right != nil {
		right = s.Right()
	}
	roff := w - len(right) - 1

	for i := 0; i < w; i++ {
		screen.SetContent(i, 0, ' ', nil, style)
	}

	for i, r := range left {
		screen.SetContent(loff+i, 0, r, nil, style)
	}

	for i, r := range center {
		screen.SetContent(coff+i, 0, r, nil, style)
	}

	for i, r := range right {
		screen.SetContent(roff+i, 0, r, nil, style)
	}

	return nil
}
