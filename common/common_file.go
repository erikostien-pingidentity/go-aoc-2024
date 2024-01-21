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
