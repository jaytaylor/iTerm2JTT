package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/jaytaylor/iTerm2JTT/pkg/iterm2"
	"github.com/jaytaylor/iTerm2JTT/pkg/site"
)

const (
	OK                 = 0
	ReadingSrcFailed   = 1
	ParsingSrcFailed   = 2
	SiteCreationFailed = 3
)

func main() {
	var r io.Reader = os.Stdin

	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			errExit(fmt.Errorf("opening %q: %s", os.Args[1], err), ReadingSrcFailed)
		}
		defer func() {
			if err = f.Close(); err != nil {
				errExit(fmt.Errorf("closing %q: %s", os.Args[1], err), ReadingSrcFailed)
			}
		}()
		r = f
	}

	tips, err := iterm2.Parse(r)
	if err != nil {
		errExit(err, ParsingSrcFailed)
	}

	fmt.Printf("got some tips! %+v\n", tips)

	if err := renderSite(tips); err != nil {
		fmt.Printf("fatal: %s\n", err)
		os.Exit(SiteCreationFailed)
	}

	os.Exit(OK)
}

func renderSite(tips iterm2.Tips) error {
	if len(tips) == 0 {
		return errors.New("site cannot be rendered with an empty list of tips")
	}

	if err := os.RemoveAll("public"); err != nil {
		return fmt.Errorf("resetting directory %q: %s", "public", err)
	}
	if err := os.MkdirAll("public", os.FileMode(int(0755))); err != nil {
		return fmt.Errorf("creating directory %q: %s", "public", err)
	}

	homeFilePath := "public/index.html"

	home, err := site.RenderHome(tips)
	if err != nil {
		return fmt.Errorf("rendering %q: %s", homeFilePath, err)
	}

	f, err := os.OpenFile(homeFilePath, os.O_CREATE|os.O_WRONLY, os.FileMode(int(0644)))
	if err != nil {
		return fmt.Errorf("opening %q: %s", homeFilePath, err)
	}
	if _, err = f.Write(home); err != nil {
		return fmt.Errorf("writing %q: %s", homeFilePath, err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("closing %q: %s", homeFilePath, err)
	}

	var (
		next     = tips[0]
		previous = tips[len(tips)-1]
	)
	fmt.Printf("previous=% v\n", previous)

	for i, tip := range tips {
		if len(tips) > i+1 {
			next = tips[i+1]
		} else {
			next = tips[0]
		}

		var tipFilePath = fmt.Sprintf("public/%v.html", tip.ID)

		page, err := site.RenderTip(tip, previous, next)
		if err != nil {
			return fmt.Errorf("rendering %q: %s", tipFilePath, err)
		}

		if f, err = os.OpenFile(tipFilePath, os.O_CREATE|os.O_WRONLY, os.FileMode(int(0644))); err != nil {
			return fmt.Errorf("opening %q: %s", tipFilePath, err)
		}
		if _, err = f.Write(page); err != nil {
			return fmt.Errorf("writing %q: %s", tipFilePath, err)
		}
		if err = f.Close(); err != nil {
			return fmt.Errorf("closing %q: %s", tipFilePath, err)
		}

		previous = tips[i]
	}

	fmt.Printf("rendered homepage + %v tips", len(tips))
	return nil
}

func errExit(what interface{}, statusCode int) {
	fmt.Fprintf(os.Stderr, "error: %s\n", what)
	os.Exit(statusCode)
}
