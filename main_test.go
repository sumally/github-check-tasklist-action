package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExec(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "testdata/pull_request_done.json",
			expected: 0,
		},
		{
			name:     "testdata/pull_request_undone.json",
			expected: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data, _ := os.ReadFile(test.name)
			assert.Equal(t, test.expected, Exec("pull_request", data))
		})
	}
}
