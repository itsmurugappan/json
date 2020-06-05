package pkg

import (
	"gotest.tools/assert"
	"testing"
)

func TestJSONDiff(t *testing.T) {
	for _, tc := range []struct {
		inputA   string
		inputB   string
		expected []string
	}{{
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		`{"c": {"a": "1", "b": "1"}, "a": "2", "b": [{"a": "2"},{"a": "3"}]}`,
		[]string{"a", "c"},
	},{
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		`{"c": {"a": "1", "b": "1"}, "a": "2", "b": [{"a": "4"},{"a": "5"}]}`,
		[]string{"a", "b", "c"},
	},{
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		nil,
	},{
		`{"a": "1", "b": ["a": "3"},{"a": "3"}], "c": {"a": "1"}}`,
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		nil,
	}} {
		act := Diff(tc.inputA, tc.inputB)
		assert.DeepEqual(t, &tc.expected, &act)
	}
}
