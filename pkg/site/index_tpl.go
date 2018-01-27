package site

import (
	"html/template"
	"io"
	"strings"

	"github.com/jaytaylor/iTerm2JTT/pkg/iterm2"

	"github.com/Masterminds/sprig"
)

const cssStyleBlock = `
        <style type="text/css">
            body {
                font-family: "Helvetica Neue", Helvetica, Helvetica, Arial, sans-serif;
                background: #000;
                color: #fff;
            }

            ul li {
                list-style-type: none;
            }

            a {
                color: #158cb8;
                text-decoration: none;
            }
            a:hover, a:active {
                color: #107b9e;
            }

            .horizontal-bar a {
                color: #60788f;
            }
            .horizontal-bar a:hover, a:active {
                color: #000;
            }

            .horizontal-bar {
                font-size: 18px;
                background: #fff;
            }

            .content {
                color: #000;
                background: #f0f0f0;
                border: 1px solid #d7d7d7;
                padding: 0 1em 1em 1em;
            }
        </style>
`

var indexTpl = template.Must(template.New("indexTpl").Funcs(sprig.FuncMap()).Parse(strings.Trim(`<!DOCTYPE html>
<html lang="en">
    <head>
        <title>Welcome to iTerm2 JTTip</title>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta name="description" content="">
        <meta name="author" content="">`+cssStyleBlock+`
    </head>
    <body>
        <h3 class="header">Welcome to iTerm2 JTTip!</h3>

        <div class="content">
            <div>
                Currently tracking {{ .tips | len }} tips.
            </div>

            <div>
                <strong>Table of Contents</strong>
            </div>

            First: <a href="0.html">#0</a>
        </div>
    </body>
</html>
`, "\r\n")))

// RenderHome renders the Homepage index.
func RenderHome(w io.Writer, tips iterm2.Tips) error {
	context := map[string]interface{}{
		"cssStyleBlock": cssStyleBlock,
		"tips":          tips,
	}
	if err := indexTpl.Execute(w, context); err != nil {
		return err
	}
	return nil
}
