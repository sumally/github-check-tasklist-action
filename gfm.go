package main

import (
	"log"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	extast "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
)

type Task struct {
	RawText   string
	IsChecked bool
}

func ExtractTaskList(source []byte) []*Task {
	gfm := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	doc := gfm.Parser().Parse(text.NewReader(source)).OwnerDocument()

	return findTasks(doc, source)
}

func findTasks(n ast.Node, source []byte) []*Task {
	tasks := make([]*Task, 0)

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if v, ok := c.(*extast.TaskCheckBox); ok {
			parent := v.Parent()

			if parent.Type() != ast.TypeBlock {
				log.Printf("parent is not base block: %+v", parent)

				continue
			}

			tasks = append(tasks, &Task{
				RawText:   extractRawText(parent, source),
				IsChecked: v.IsChecked,
			})
		}

		tasks = append(tasks, findTasks(c, source)...)
	}

	return tasks
}

func extractRawText(n ast.Node, source []byte) string {
	sb := strings.Builder{}

	for i := 0; i < n.Lines().Len(); i++ {
		t := n.Lines().At(i)
		sb.Write(t.Value(source))
	}

	return sb.String()
}
