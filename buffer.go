package vogue

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/gdamore/tcell"
)

type Buffer struct {
	Path string
	Name string
	Data string
	Line int
	Col  int
}

var Empties int = 0

var (
	BadPath   error = errors.New("bad path")
	Irregular error = errors.New("not a regular file")
)

func BufEmpty() *Buffer {
	Empties++
	return &Buffer{
		Path: "",
		Name: "empty #" + strconv.Itoa(Empties),
		Data: "",
		Line: 0,
		Col:  0,
	}
}

func BufFile(path string) (*Buffer, error) {
	b := &Buffer{}

	b.Path = path

	parts := strings.Split(b.Path, string([]rune{os.PathSeparator}))
	if len(parts) < 1 {
		return nil, BadPath
	}
	b.Name = parts[len(parts)-1]

	info, err := os.Stat(b.Path)
	if err != nil {
		return nil, err
	}
	if info.Mode()&os.ModeType != 0 {
		return nil, Irregular
	}

	f, err := os.Open(b.Path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	b.Data = string(data)

	b.Line = 0
	b.Col = 0

	return b, nil
}

func (b *Buffer) Key(evk *tcell.EventKey) {
	panic("buffer key stub")
}
