package common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
)

func GetLinesFromTestdataFile(filename string) []string {
	// Get the file path of the stack frame that called this function
	_, callerFilepath, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("error: no caller information"))
	}

	callerFileDir := path.Dir(callerFilepath)

	// Get the full pathname for the data
	dataFilePath := fmt.Sprintf("%s/testdata/%s", callerFileDir, filename)

	// Open the puzzle data that implements the io.Reader interface
	dataFile, err := os.Open(dataFilePath)
	if err != nil {
		panic(err)
	}

	// Make sure to close the opened file once this scope is closed
	defer dataFile.Close()

	// Create a new Scanner to enable reading the file line by line
	scanner := bufio.NewScanner(dataFile)

	// Read each line and populate string splice to return
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// Generic Prepend function to place a value at the 0th index of a slice
func Prepend[K any](genericSlice []K, prependData K) []K {
	// Add the prepend data to the end of the slice
	// This is to extend the slice length by one
	// The data will be shifted destructively in the next step
	genericSlice = append(genericSlice, prependData)

	// Shift data destructively to the right
	copy(genericSlice[1:], genericSlice)

	// Append needed data to front of slice
	genericSlice[0] = prependData

	return genericSlice
}
