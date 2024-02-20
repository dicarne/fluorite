package main

import (
	"bytes"
	"path"
	"strings"
)

func wrapHTML(html []byte, filename string, style *StyleConfig, baseUrl string, isIndex bool) []byte {
	buf := bytes.Buffer{}
	buf.WriteString("<html lang=\"")
	buf.WriteString(*lang)
	buf.WriteString("\">")
	buf.WriteString(`
		<head>
			<meta charset="utf-8">`)
	if index_web_path != "" && isIndex {
		buf.WriteString(`
			<meta http-equiv="refresh" content="0;url=notes/`)
		buf.WriteString(index_web_path)
		buf.WriteString(`">`)
	}
	for i := range style.Css {
		_writeStyle(&buf, _makeUrl(baseUrl, style.Name, style.Css[i]))
	}
	for i := range style.Js {
		_writeScript(&buf, _makeUrl(baseUrl, style.Name, style.Js[i]))
	}
	buf.WriteString(`<title>`)
	buf.WriteString(_shortTitle(filename))
	buf.WriteString(`
		</title>
		</head>`)
	buf.WriteString("<body>")
	buf.WriteString("<div class=\"note-body\" id=\"note-body\">")
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

func _shortTitle(title string) string {
	paths := strings.Split(strings.Split(title, ".")[0], "/")
	return paths[len(paths)-1]
}
