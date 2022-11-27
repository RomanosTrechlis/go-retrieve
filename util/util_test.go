package util_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/RomanosTrechlis/go-retrieve/util"
)

func TestCSVToArray(t *testing.T) {
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

func TestMarshalIndent(t *testing.T) {
	s := struct {
		name string
	}{
		name: "Name",
	}
	_, err := util.MarshalIndent(s, false)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = util.MarshalIndent(s, true)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestLoadFile(t *testing.T) {
	_, err := util.LoadFile("util.go")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	b, err := util.LoadFile("not_existing.txt")
	if err == nil {
		t.Fatalf("expected error, got %v", string(b))
	}
}

func TestWriteFile(t *testing.T) {
	testContent := struct {
		field  string
		isTest bool
	}{
		field:  "test",
		isTest: true,
	}
	defer os.RemoveAll("data")
	defer os.RemoveAll("temp_data")

	err := util.WriteFile("data", "test.json", testContent)
	if err != nil {
		t.Errorf("failed to write file: %v", err)
	}
	err = util.WriteFile("temp_data", "test.json", testContent)
	if err != nil {
		t.Errorf("failed to write file: %v", err)
	}
	err = util.WriteFile("temp_data", "test.yml", nil)
	if err != nil {
		t.Errorf("failed to write file: %v", err)
	}
}

func TestContains(t *testing.T) {
	testCases := []struct {
		description string
		inputArray  []string
		inputString string
		output      bool
	}{
		{"contains", []string{"a", "b", "c"}, "c", true},
		{"doesn't contain", []string{"a", "b", "c"}, "d", false},
		{"empty array", []string{}, "d", false},
		{"nil array", nil, "d", false},
		{"empty string", []string{"a", "b", "c"}, "", false},
		{"complex", []string{"a", "b", "c", "d\ne"}, "d\ne", true},
	}

	for _, test := range testCases {
		contains := util.Contains(test.inputArray, test.inputString)
		if contains != test.output {
			t.Errorf("failed %s test: expected for %v, got %v", test.description, test.output, contains)
		}
	}
}

func TestIsExists(t *testing.T) {
	exists := util.IsExists("nonExisting.json")
	if exists {
		t.Errorf("failed: expected for nonExisting.json not to exists, but it does")
	}

	exists = util.IsExists(filepath.Join("..", "cli", "data", "config.json"))
	if !exists {
		t.Errorf("failed: expected for ../cli/data/config.json to exists, but it doesn't")
	}
}

type scanTestCase struct {
	description string
	input       string
	output      string
}

func TestScan(t *testing.T) {
	testCases := []scanTestCase{
		{"simple string input", "test", "test"},
		{"complex string input", "this is a test", "this is a test"},
		{"new line string input", "\n", ""},
	}

	for _, test := range testCases {
		scanTest(t, test)
	}
}

func scanTest(t *testing.T, test scanTestCase) {
	content := []byte(test.input)
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Errorf("failed to create temp file")
	}

	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Errorf("failed to create temp file")
	}

	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Errorf("failed to create temp file")
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin

	os.Stdin = tmpfile
	s := util.Scan("test prompt")
	if s != test.output {
		t.Errorf("failed %s: expected %s, got %s", test.description, test.output, s)
	}

	_ = tmpfile.Close()
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
