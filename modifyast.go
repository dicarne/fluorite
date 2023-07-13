package main

import (
	"bytes"

	"github.com/gomarkdown/markdown/ast"
)

func modifyAst(doc ast.Node) ast.Node {
	doclinkPrefixBytes := []byte("[[")
	doclinkSuffixBytes := []byte("]]")
	doclinkPrefixBytesLen := len(doclinkPrefixBytes)
	ast.WalkFunc(doc, func(node ast.Node, entering bool) ast.WalkStatus {
		if entering {
			leaf := node.AsLeaf()
			if leaf != nil {

				doclinkPrefixIndex := bytes.Index(leaf.Literal, doclinkPrefixBytes)
				if doclinkPrefixIndex == -1 {
					return ast.GoToNext
				}
				doclinkSuffixIndex := bytes.Index(leaf.Literal[doclinkPrefixIndex:], doclinkSuffixBytes)
				if doclinkSuffixIndex == -1 {
					return ast.GoToNext
				}
				if doclinkPrefixIndex > doclinkSuffixIndex {
					return ast.GoToNext
				}
				if doclinkPrefixIndex != 0 {
					newnode := &ast.Text{}
					newnode.Literal = leaf.Literal[:doclinkPrefixIndex]
					leaf.Literal = leaf.Literal[doclinkPrefixIndex:]
					index := 0
					parent := node.GetParent().AsContainer()
					for i, v := range parent.Children {
						if v == node {
							index = i
							break
						}
					}
					insertNewNode(newnode, parent, index)

					return ast.GoToNext
				}

				insideDoclink := leaf.Literal[doclinkPrefixIndex+doclinkPrefixBytesLen : doclinkSuffixIndex]

				leaf.Literal = leaf.Literal[doclinkSuffixIndex+doclinkPrefixBytesLen:]

				parent := node.GetParent().AsContainer()
				index := 0
				for i, v := range parent.Children {
					if v == node {
						index = i
						break
					}
				}
				newnode := &DocLink{
					URL:    string(insideDoclink),
					Inline: true,
				}
				insertNewNode(newnode, parent, index)

				// debug
				// fmt.Println("--")
				// for _, v := range parent.Children {
				// 	vv := v.AsLeaf()
				// 	if vv != nil {
				// 		fmt.Println(string(vv.Literal))
				// 	}
				// }
			}
		}
		return ast.GoToNext
	})
	return doc
}

func insertNewNode(newnode ast.Node, parentn ast.Node, at int) {
	newnode.SetParent(parentn)
	parent := parentn.AsContainer()
	pc := parent.Children
	tmp := append([]ast.Node{}, pc[at:]...)
	pc = append(pc[0:at], newnode)
	pc = append(pc, tmp...)
	parent.SetChildren(pc)
}
