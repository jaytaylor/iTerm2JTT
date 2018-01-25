package site

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/jaytaylor/iTerm2JTT/pkg/iterm2"

	"github.com/Masterminds/sprig"
)

var indexTpl = template.Must(template.New("indexTpl").Funcs(sprig.FuncMap()).Parse(strings.Trim(`<!DOCTYPE html>
<html lang="en">
	<head>
	    <title>Welcome to iTerm2 JTTip</title>
	    <meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta name="description" content="">
		<meta name="author" content="">
	</head>
	<body>
		<h3 id="header">Welcome to iTerm2 JTTip!</h3>

		<div id="content">
			Currently tracking {{ .tips | len }} tips.

			First: <a href="0.html">#0</a>
		</div>
	</body>
</html>
`, "\r\n")))

// RenderHome renders the Homepage index.
func RenderHome(tips iterm2.Tips) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := indexTpl.Execute(buf, map[string]interface{}{"tips": tips}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
