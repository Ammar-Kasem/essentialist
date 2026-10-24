package main

import (
	"io"
	"strings"

	terminal "github.com/MichaelMure/go-term-markdown"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

type flashdownRenderer struct {
	markdown.Renderer
}

func (r flashdownRenderer) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	// Handle parsed HTML spans so tags inside code remain literal text.
	if span, ok := node.(*ast.HTMLSpan); ok {
		switch strings.ToLower(string(span.Literal)) {
		case "<br>", "<br/>", "<br />":
			return r.Renderer.RenderNode(w, &ast.Hardbreak{}, entering)
		}
	}
	return r.Renderer.RenderNode(w, node, entering)
}

func renderMarkdown(source string, width int) []byte {
	p := parser.NewWithExtensions(terminal.Extensions())
	nodes := markdown.Parse([]byte(source), p)
	r := flashdownRenderer{Renderer: terminal.NewRenderer(width, 0)}
	return markdown.Render(nodes, r)
}
