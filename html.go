package main

import (
	"bytes"
	"path"
	"strings"
)

func wrapHTML(html []byte, filename string, style *StyleConfig, baseUrl string) []byte {
	buf := bytes.Buffer{}
	buf.WriteString("<html>")
	buf.WriteString(`
		<head>
		<meta charset="utf-8">`)
	for i := range style.Css {
		_writeStyle(&buf, _makeUrl(baseUrl, style.Name, style.Css[i]))
	}
	for i := range style.Js {
		_writeScript(&buf, _makeUrl(baseUrl, style.Name, style.Js[i]))
	}
	buf.WriteString(`<title>`)
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

func _writeStyle(buff *bytes.Buffer, path string) {
	buff.WriteString(`<link rel="stylesheet" type="text/css" href="`)
	buff.WriteString(path)
	buff.WriteString(`">`)
}

func _writeScript(buff *bytes.Buffer, path string) {
	buff.WriteString(`<script src="`)
	buff.WriteString(path)
	buff.WriteString(`""></script>`)
}

func _makeUrl(base string, theme string, url string) string {
	if strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://") {
		return url
	}
	return path.Join(base, "theme", theme, url)
}
