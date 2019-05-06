package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/qpliu/qrencode-go/qrencode"
	"io/ioutil"
	"os"
)

type cmdOptions struct {
	Help    bool `short:"h" long:"help" description:"show this help message"`
}

func showHelp() {
	const v = `Usage: qrcode [OPTIONS] [TEXT]

Options:
 -h, --help
   Show this help message
`

	os.Stderr.Write([]byte(v))
}

func pErr(format string, a ...interface{}) {
	fmt.Fprint(os.Stdout, os.Args[0], ": ")
	fmt.Fprintf(os.Stdout, format, a...)
}

func main() {
	ret := 0
	defer func() { os.Exit(ret) }()

	opts := &cmdOptions{}
	optsParser := flags.NewParser(opts, flags.PrintErrors)
	args, err := optsParser.Parse()
	if err != nil || len(args) > 1 {
		showHelp()
		ret = 1
		return
	}
	if opts.Help {
		showHelp()
		return
	}

	var text string
	if len(args) == 1 {
		text = args[0]
	} else {
		text_bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			pErr("read from stdin failed: %v\n", err)
			ret = 1
			return
		}
		text = string(text_bytes)
	}

	grid, err := qrencode.Encode(text, qrencode.ECLevelL)
	if err != nil {
		pErr("encode failed: %v\n", err)
		ret = 1
		return
	}

	height := grid.Height()
	width := grid.Width()

	WHITE_ALL := '\u2588'
	WHITE_BLACK := '\u2580'
	BLACK_WHITE := '\u2584'
	BLACK_ALL := ' '

	for x := 0; x <= width + 1; x++ {
		fmt.Printf("%c", BLACK_WHITE)
	}

	fmt.Printf("\n")

	for y := 0; y < height; y += 2 {

		fmt.Printf("%c", WHITE_ALL)

		for x := 0; x < width; x++ {

			currentIsBlack := grid.Get(x,y)

			nextIsBlack := false;
			if (y + 1 < height) {
				nextIsBlack = grid.Get(x, y+1)
			}

			if currentIsBlack && nextIsBlack {
				fmt.Printf("%c", BLACK_ALL)
			} else if currentIsBlack && !nextIsBlack {
				fmt.Printf("%c", BLACK_WHITE)
			} else if !currentIsBlack && nextIsBlack {
				fmt.Printf("%c", WHITE_BLACK)
			} else if !currentIsBlack && !nextIsBlack {
				fmt.Printf("%c", WHITE_ALL)
			}
		}

		fmt.Printf("%c\n", WHITE_ALL)

	}

}
