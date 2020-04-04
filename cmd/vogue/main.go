package main

import (
	"flag"

	"github.com/superloach/vogue"
)

func main() {
	flag.Parse()
	args := flag.Args()

	v, err := vogue.NewVogue()
	if err != nil {
		panic(err)
	}
	defer v.Fini()

	if len(args) == 0 {
		v.Tabs.Buffers = append(
			v.Tabs.Buffers,
			vogue.BufEmpty(),
		)
	}

	for _, arg := range args {
		buf, err := vogue.BufFile(arg)
		if err != nil {
			panic(err)
		}

		v.Tabs.Buffers = append(v.Tabs.Buffers, buf)
	}

	v.Fresh()

	err = v.Run()
	if err != nil {
		panic(err)
	}
}
