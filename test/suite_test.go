package test

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertFile(t *testing.T, size int, fileName string, isRemove bool) {
	file, _ := os.Open(fileName)

	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}

	defer file.Close()

	assert.FileExists(t, fileName)
	assert.Equal(t, size, lineCount)

	if isRemove {
		os.Remove(fileName)
	}
}
