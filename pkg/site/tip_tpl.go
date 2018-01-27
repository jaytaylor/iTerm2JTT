package site

import (
	"html/template"
	"io"
	"strings"

	"github.com/jaytaylor/iTerm2JTT/pkg/iterm2"

	"github.com/Masterminds/sprig"
)

var tipTpl = template.Must(template.New("tipTpl").Funcs(sprig.FuncMap()).Parse(strings.Trim(`<!DOCTYPE html>
<html lang="en">
    <head>
        <title>iTerm2 JTTip | {{ .tip.Title }} </title>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta name="description" content="">
        <meta name="author" content="">`+cssStyleBlock+`
    </head>
    <body>

        <h3 class="header">{{ .tip.Title }}</h3>

        #{{ .tip.ID }}

        <div class="content">
            <div class="horizontal-bar">
                <ul>
                    <li><a href="index.html">Home</a></li>
                </ul>
            </div>

            {{ .tip.Body }}

            {{ if gt (len .tip.URL) 0 }}
            <div id="more">
                <a href="{{ .tip.URL }}">Read More</a>
            </div>
            {{ end }}
        </div>

        {{ .previous.Title }}

        <a href="{{ .previous.ID }}.html">Previous ({{ .previous.Title }})</a> | <a href="{{ add .next.ID 1 }}.html">Next ({{ .next.Title }})</a>
    </body>
</html>
`, "\r\n")))

// RenderTip converts a Tip into a site HTML page content.
func RenderTip(w io.Writer, tip iterm2.Tip, previous iterm2.Tip, next iterm2.Tip) error {
	context := map[string]interface{}{
		"tip":      tip,
		"previous": previous,
		"next":     next,
	}

	if err := tipTpl.Execute(w, context); err != nil {
		return err
	}
	return nil
}
