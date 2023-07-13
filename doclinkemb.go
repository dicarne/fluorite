package main

import (
	"bytes"

	"github.com/gomarkdown/markdown/ast"
)

var doclinkemb_prefix = []byte("![[")
var doclinkemb_suffix = []byte("]]")

type DocLinkEmb struct {
	ast.Leaf
	URL string
}

func ParseDocLinkEmb(data []byte) (ast.Node, []byte, int) {
	if !bytes.HasPrefix(data, doclinkemb_prefix) {
		return nil, nil, 0
	}
	suffixLen := len(doclinkemb_suffix)
	i := len(doclinkemb_prefix)
	end := bytes.Index(data[i:], doclinkemb_suffix)
	if end < 0 {
		return nil, data, 0
	}
	end = end + i
	lines := string(data[i:end])
	if len(data) > end {
		if data[end+suffixLen] != '\n' {
			return nil, data, 0
		}
	}
	res := &DocLinkEmb{
		URL: lines,
	}

	return res, nil, end + suffixLen
}
