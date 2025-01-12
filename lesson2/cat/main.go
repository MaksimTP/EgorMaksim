package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	B bool
	E bool
	N bool
	S bool
	T bool
}

func ParseArgs() (Options, []string) {
	b := flag.Bool("b", false, "")
	n := flag.Bool("n", false, "")
	e := flag.Bool("e", false, "")
	s := flag.Bool("s", false, "")
	t := flag.Bool("t", false, "")
	flag.Parse()

	files := flag.Args()
	return Options{
		B: *b,
		E: *e,
		N: *n,
		S: *s,
		T: *t,
	}, files
}

func Cat(filename string, opts Options) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	rows := make([]string, 0)
	row := ""
	for _, char := range data {
		if char == 10 {
			row += string(char)
			rows = append(rows, row)
			row = ""
		} else {
			row += string(char)
		}
	}

	for numRow, row := range rows {
		if opts.N {
			fmt.Printf("\t%d ", numRow+1)
		}

		fmt.Print(row)
	}
}

func main() {
	opts, files := ParseArgs()
	for _, file := range files {
		Cat(file, opts)
	}

}
