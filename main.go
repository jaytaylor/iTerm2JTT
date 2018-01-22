package main

import (
	"fmt"
	"io"
	"os"

	"github.com/jaytaylor/iTerm2JTT/pkg/iterm2"
)

func main() {
	var r io.Reader = os.Stdin

	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			errExit(fmt.Errorf("opening file %q: %s", os.Args[1], err), 2)
		}
		defer func() {
			if err = f.Close(); err != nil {
				errExit(fmt.Errorf("closing file %q: %s", os.Args[1], err), 3)
			}
		}()
		r = f
	}

	tips, err := iterm2.Parse(r)
	if err != nil {
		errExit(err, 1)
	}

	fmt.Printf("got some tips! %+v\n", tips)
}

func errExit(what interface{}, statusCode int) {
	fmt.Fprintf(os.Stderr, "error: %s\n", what)
}
