package vogue

import "github.com/gdamore/tcell"

type Tabs struct {
	Current int
	Buffers []*Buffer
}

func NewTabs() *Tabs {
	return &Tabs{
		Current: 0,
		Buffers: make([]*Buffer, 0),
	}
}

func (t *Tabs) Buffer() *Buffer {
	return t.Buffers[t.Current]
}

func (t *Tabs) Key(evk *tcell.EventKey) {
	panic("buffer key stub")
}
