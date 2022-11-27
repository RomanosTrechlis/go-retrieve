package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"strings"
)

// WriteFile creates a file with the given filename at the given
// filePath containing the json data passed in the parameter.
func WriteFile(filePath, filename string, data interface{}) error {
	file, err := MarshalIndent(data, false)
	if err != nil {
		return err
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	configFile := path.Join(filePath, filename)
	return os.WriteFile(configFile, file, 0755)
}

// LoadFile returns the bytes of a file if it exists
func LoadFile(filename string) ([]byte, error) {
	if !IsExists(filename) {
		return nil, fmt.Errorf("couldn't find file")
	}

	return os.ReadFile(filename)
}

// Scan prints a prompt and wait for user input to return
func Scan(prompt string) string {
	prompt = strings.TrimRight(prompt, " ")
	fmt.Printf(prompt + " ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// IsExists checks for the existence of the given file or directory
func IsExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// Contains checks a string array for a specific string
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// CSVToArray breaks a csv string to a string array
func CSVToArray(csv string) []string {
	if strings.Trim(csv, " ") == "" {
		return []string{}
	}
	arr := strings.Split(csv, ",")
	for i := 0; i < len(arr); i++ {
		arr[i] = strings.Trim(arr[i], " ")
	}
	return arr
}

func MarshalIndent(c interface{}, isJson bool) ([]byte, error) {
	if isJson {
		return json.MarshalIndent(c, "", "  ")
	}
	return yaml.Marshal(c)
}
