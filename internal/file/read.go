package file

import (
	"bufio"
	"os"
)

// Read return the slice of string of all the line in the given file
func Read(filePath string) (txtlines []string, file *os.File) {
	// open the file
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_SYNC, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// read the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// add the lines to the varible from the file
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	return
}
