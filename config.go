package vogue

import (
	"encoding/json"
	"os"

	"github.com/gdamore/tcell"
)

type Config map[string]interface{}

func ConfigFile(flag int, perm os.FileMode) (*os.File, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	sep := string([]rune{os.PathSeparator})

	path += sep + "vogue"

	err = os.MkdirAll(path, os.ModeDir|perm|0100)
	if err != nil {
		return nil, err
	}

	path += sep + "config.json"

	return os.OpenFile(path, flag, perm)
}

func LoadConfig() (Config, error) {
	f, err := ConfigFile(os.O_RDONLY, 0744)
	if err != nil {
		return nil, err
	}

	c := make(Config)
	err = json.NewDecoder(f).Decode(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c Config) Save() error {
	f, err := ConfigFile(os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		return err
	}

	return json.NewEncoder(f).Encode(&c)
}

func (c Config) Get(name string, def interface{}) interface{} {
	item, ok := c[name]
	if !ok {
		c[name] = def
		return def
	}
	return item
}

func (c Config) GetStyle(name string, def tcell.Style) tcell.Style {
	istyle := c.Get("style."+name, float64(def))

	style, ok := istyle.(float64)
	if !ok {
		return def
	}

	return tcell.Style(style)
}
