package site

import (
	"bytes"
	"html/template"
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
	<meta name="author" content="">
</head>
<body>

<a href="/">Home</a>

<h3 id="header">{{ .tip.Title }}</h3>

#{{ .tip.ID }}

<div id="content">
{{ .tip.Body }}
</div>

{{ if gt (len .tip.URL) 0 }}
<div id="more">
    <a href="{{ .tip.URL }}">Read More</a>
</div>
{{ end }}

{{ .previous.Title }}

<a href="{{ .previous.ID }}.html">Previous</a> | <a href="{{ add .next.ID 1 }}.html">Next</a>
</body>
</html>
`, "\r\n")))

// RenderTip converts a Tip into a site HTML page content.
func RenderTip(tip iterm2.Tip, previous iterm2.Tip, next iterm2.Tip) ([]byte, error) {
	var (
		context = map[string]interface{}{
			"tip":      tip,
			"previous": previous,
			"next":     next,
		}
		buf = &bytes.Buffer{}
	)

	if err := tipTpl.Execute(buf, context); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
