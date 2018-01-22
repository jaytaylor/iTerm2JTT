package iterm2

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

// Parse consumes the contents of an iTermTipData.m file and produces a
// populated Tips object.
func Parse(r io.Reader) (Tips, error) {
	var (
		startExpr = regexp.MustCompile(`@[ \t]*"([0-9]{3,})"`)
		fieldExpr = regexp.MustCompile(`kTip(Title|Body|Url)Key[ \t]*:[ \t]*@[ \t]*"(.*)"`)
		scanner   = bufio.NewScanner(r)
		tips      = Tips{}
		tip       *Tip
		line      string
	)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		line = scanner.Text()
		if matches := startExpr.FindStringSubmatch(line); len(matches) > 0 {
			if tip != nil && !tip.Empty() {
				tips = append(tips, *tip)
			}

			id, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, fmt.Errorf("parsing id %q: %s", matches[1], err)
			}
			tip = &Tip{
				ID: id,
			}
		}

		if matches := fieldExpr.FindStringSubmatch(line); len(matches) > 0 {
			if tip == nil {
				tip = &Tip{}
			}
			switch matches[1] {
			case "Title":
				tip.Title = matches[2]
			case "Body":
				tip.Body = matches[2]
			case "Url":
				tip.URL = matches[2]
			}
		}
	}

	if tip != nil && !tip.Empty() {
		tips = append(tips, *tip)
	}

	return tips, nil
}
