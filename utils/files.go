package utils

import (
	"os"
	"slices"
	"strings"
)

func GetFileString(path string) string {
	// Load the data
	dat, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func GetFileLines(path string) []string {
	// Load the file as one string with \n
	file := GetFileString(path)
	// Split the string into lines
	lines := strings.Split(file, "\n")
	// Remove empty lines
	lines = slices.DeleteFunc(lines, func(s string) bool {
		return s == ""
	})
	return lines
}
