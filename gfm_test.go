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

// A task list marker is only a task when it starts the list item, so a checkbox
// that follows other text is ordinary content. goldmark before v1.8.5 reported
// those as tasks, which made a body such as "- foo [ ] bar" fail the check.
func TestExtractTaskListOnlyLeadingCheckbox(t *testing.T) {
	tests := []struct {
		name     string
		md       string
		expected []*Task
	}{
		{
			name:     "a checkbox after text is not a task",
			md:       "- foo [ ] bar\n",
			expected: []*Task{},
		},
		{
			name: "only the leading checkbox is a task",
			md:   "- [x] a [ ] b\n",
			expected: []*Task{
				{
					RawText:   "[x] a [ ] b",
					IsChecked: true,
				},
			},
		},
		{
			name: "a repeated checkbox belongs to the text",
			md:   "- [ ] [ ] double\n",
			expected: []*Task{
				{
					RawText:   "[ ] [ ] double",
					IsChecked: false,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, ExtractTaskList([]byte(test.md)))
		})
	}
}
