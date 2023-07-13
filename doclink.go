package main

import (
	"bytes"

	"github.com/gomarkdown/markdown/ast"
)

var doclink_prefix = []byte("[[")
var doclink_suffix = []byte("]]")

type DocLink struct {
	ast.Leaf
	URL    string
	Inline bool
}

func ParseDocLink(data []byte) (ast.Node, []byte, int) {
	if !bytes.HasPrefix(data, doclink_prefix) {
		return nil, nil, 0
	}
	suffixLen := len(doclinkemb_suffix)
	i := len(doclink_prefix)
	end := bytes.Index(data[i:], doclink_suffix)
	if end < 0 {
		return nil, data, 0
	}
	end = end + i

	lines := string(data[i:end])

	if len(data) > end && len(data) > end+suffixLen {
		if data[end+suffixLen] != '\n' {
			return nil, data, 0
		}
	}
	res := &DocLink{
		URL:    lines,
		Inline: false,
	}
	return res, nil, end + suffixLen
}
