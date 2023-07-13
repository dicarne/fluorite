package main

import (
	"bytes"
)

func wrapHTML(html []byte, filename string, css string) []byte {
	buf := bytes.Buffer{}
	buf.WriteString("<html>")
	buf.WriteString(`
		<head>
		<meta charset="utf-8">
		<link rel="stylesheet" type="text/css" href="`)
	buf.WriteString(css)
	buf.WriteString(`"><title>`)
	buf.WriteString(filename)
	buf.WriteString(`
		</title>
		</head>`)
	buf.WriteString("<body>")
	buf.WriteString("<div class=\"note-body\">")
	buf.Write(html)
	buf.WriteString("</div>")
	buf.WriteString("</body>")
	buf.WriteString("</html>")
	return buf.Bytes()
}
