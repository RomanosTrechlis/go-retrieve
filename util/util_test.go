package util_test

import (
	"github.com/RomanosTrechlis/go-retrieve/util"
	"testing"
)

func Test_CSVToArray(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		output      []string
	}{
		{description: "empty input", input: "", output: []string{}},
		{description: "one csv value", input: "test", output: []string{"test"}},
		{description: "multiple csv values", input: "test,big,use", output: []string{"test", "big", "use"}},
		{description: "csv values with spaces", input: "test, big, use", output: []string{"test", "big", "use"}},
		{description: "csv values with spaces on string values", input: " test , big , use ", output: []string{"test", "big", "use"}},
	}

	for _, test := range testCases {
		o := util.CSVToArray(test.input)
		if !equal(o, test.output) {
			t.Errorf("test for '%s' failed: expected '%v', got '%v'", test.description, test.output, o)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
