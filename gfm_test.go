package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractTaskList(t *testing.T) {
	md := `
# foo

## bar

- [x] todo 1
  - [ ] *todo* 1-1
  - [x] todo 1-2
- [ ] todo 2

- [] invalid

` + "```\n- [ ] ignore\n```"

	t.Log(md)

	assert.Equal(t, []*Task{
		{
			RawText:   "[x] todo 1",
			IsChecked: true,
		},
		{
			RawText:   "[ ] *todo* 1-1",
			IsChecked: false,
		},
		{
			RawText:   "[x] todo 1-2",
			IsChecked: true,
		},
		{
			RawText:   "[ ] todo 2",
			IsChecked: false,
		},
	}, ExtractTaskList([]byte(md)))
}
