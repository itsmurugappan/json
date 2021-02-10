package transform

import (
	"testing"

	"gotest.tools/assert"
)

func TestJSONDiff(t *testing.T) {
	for _, tc := range []struct {
		name     string
		inputA   string
		inputB   string
		expected []string
	}{{
		"multi level mismatch case 1",
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		`{"c": {"a": "1", "b": "1"}, "a": "2", "b": [{"a": "2"},{"a": "3"}]}`,
		[]string{"a", "c"},
	}, {
		"multi level mismatch case 2",
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		`{"c": {"a": "1", "b": "1"}, "a": "2", "b": [{"a": "4"},{"a": "5"}]}`,
		[]string{"a", "b", "c"},
	}, {
		"no mismatch",
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		nil,
	}, {
		"wrong json",
		`{"a": "1", "b": ["a": "3"},{"a": "3"}], "c": {"a": "1"}}`,
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		nil,
	}} {
		t.Run(tc.name, func(t *testing.T) {
			act := Diff(tc.inputA, tc.inputB)
			assert.DeepEqual(t, &tc.expected, &act)
		})
	}
}

func TestObjectDiff(t *testing.T) {
	for _, tc := range []struct {
		name       string
		inputA     string
		inputB     string
		inputAPath string
		inputBPath string
		expected   []string
	}{{
		"first level match",
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		`{"c": {"a": "1", "b": "1"}, "a": "2", "b": [{"a": "2"},{"a": "3"}]}`,
		"b",
		"b",
		nil,
	}, {
		"first level mismatch",
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`,
		`{"c": {"a": "1", "b": "1"}, "a": "2", "b": [{"a": "2"},{"a": "3"}]}`,
		"c",
		"c",
		[]string{"b"},
	}, {
		"second level mismatch",
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": {"ab": 1, "cd": 2}}}`,
		`{"c": {"a": {"ab": 1, "cd": 3}, "b": "1"}, "a": "2", "b": [{"a": "2"},{"a": "3"}]}`,
		"c.a",
		"c.a",
		[]string{"cd"},
	}, {
		"object missing",
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": {"ab": 1, "cd": 2}}}`,
		`{"c": {"b": "1"}, "a": "2", "b": [{"a": "2"},{"a": "3"}]}`,
		"c.a",
		"c.a",
		nil,
	}, {
		"second level match",
		`{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": {"ab": 1, "cd": 2}}}`,
		`{"c": {"e": {"ab": 1, "cd": 2}, "b": "1"}, "a": "2", "b": [{"a": "2"},{"a": "3"}]}`,
		"c.a",
		"c.e",
		nil,
	}} {
		t.Run(tc.name, func(t *testing.T) {
			act := ObjectDiff(tc.inputA, tc.inputB, tc.inputAPath, tc.inputBPath)
			assert.DeepEqual(t, &tc.expected, &act)
		})
	}
}
