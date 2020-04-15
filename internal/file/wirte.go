package file

import (
	"fmt"
	"os"
)

// Write to the file and close the connection
func Write(t string, f *os.File) {
	// write to the file from the beginning
	l, err := f.WriteAt([]byte(t), 0)
	if err != nil {
		f.Close()
		panic(err)
	}
	fmt.Printf("> %v bytes written: ", l)
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
