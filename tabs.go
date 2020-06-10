package vogue

import (
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
)

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
	print("tabs key stub")
}

func (t *Tabs) Fresh(screen tcell.Screen, focus tcell.Style, style tcell.Style) {
	cur := t.Buffer()

	befores := []string{}
	for i := 0; i < t.Current; i++ {
		befores = append(befores, strconv.Itoa(i+1)+") "+t.Buffers[i].Name)
	}
	before := []rune(strings.Join(befores, " : "))

	current := []rune(" | " + cur.Name + " | ")

	afters := []string{}
	for i := t.Current + 1; i < len(t.Buffers); i++ {
		afters = append(afters, strconv.Itoa(i+1)+") "+t.Buffers[i].Name)
	}
	after := []rune(strings.Join(afters, " : "))

	w, _ := screen.Size()

	coff := (w / 2) - (len(current) / 2)
	boff := coff - len(before)
	aoff := coff + len(current)

	for i := 0; i < len(before); i++ {
		screen.SetContent(boff + i, 1, before[i], nil, style)
	}

	for i := 0; i < len(current); i++ {
		screen.SetContent(coff + i, 1, current[i], nil, style)
	}

	for i := 0; i < len(after); i++ {
		screen.SetContent(aoff + i, 1, after[i], nil, style)
	}
}
