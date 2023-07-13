package main

import (
	"bytes"
	"strings"

	"github.com/gomarkdown/markdown/ast"
)

type Callout struct {
	ast.Container
	Tag   string
	Title string
}

func ParseCallout(data []byte) (ast.Node, []byte, int) {
	var prefix = []byte("> [!")
	var suffix = []byte("]")

	if !bytes.HasPrefix(data, prefix) {
		return nil, nil, 0
	}
	i := len(prefix)
	tagend := bytes.Index(data[i:], suffix)
	if tagend < 0 {
		return nil, data, 0
	}
	tagend = tagend + i
	tag := string(data[i:tagend])
	titleStart := bytes.Index(data[tagend:], []byte(" "))
	titleEnd := bytes.Index(data[tagend:], []byte("\n"))

	title := ""
	if titleStart < titleEnd {
		title = string(data[titleStart+tagend : titleEnd+tagend])
	}

	res := &Callout{
		Tag:   tag,
		Title: title,
	}

	blockStart := bytes.Index(data[i:], []byte("> "))
	finalEnd := bytes.Index(data[i:], []byte("\n\n"))

	lines := strings.Split(string(data[i+blockStart:finalEnd+i]), "> ")
	blockbuffer := []byte(strings.Join(lines, "\n"))

	return res, blockbuffer, finalEnd + 2 + i
}
