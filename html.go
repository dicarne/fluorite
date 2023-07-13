package main

import (
	"bytes"
)

func wrapHTML(html []byte, filename string) []byte {
	buf := bytes.Buffer{}
	buf.WriteString("<html>")
	buf.WriteString(`
		<head>
		<meta charset="utf-8">
		<link rel="stylesheet" type="text/css" href="../main.css">
		<title>`)
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
