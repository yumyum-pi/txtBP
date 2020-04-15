package internal

import (
	"io/ioutil"
	"path/filepath"
)

// getPath the list of file path
func getPath(d, e string) (fPath []string) {
	// get the slice of files in the given directory
	fs, err := ioutil.ReadDir(d)
	if err != nil {
		panic(err)
	}

	// loop through all the files
	for _, file := range fs {
		// check for the extension
		if filepath.Ext(file.Name()) == e {

			p := filepath.Join(d, file.Name()) // join root directory and file name
			fPath = append(fPath, p)           // add the path to the slice
		}
	}

	return
}
