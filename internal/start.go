package internal

import (
	"fmt"
	"os"

	"github.com/yumyum-pi/txtBP/internal/file"
)

// Start the process of the
func Start(d, e, s string) {
	var fPaths []string     // store the list of file paths
	fPaths = []string{d}    // set the defult to the given path
	info, err := os.Stat(d) // get path info

	// the path does not  exist
	if os.IsNotExist(err) {
		fmt.Println("> Error:", d, " path not found")
		os.Exit(1) // exit the program
	}

	// check if the path is a direcotry
	if info.IsDir() {
		fPaths = getPath(d, e) // get the file paths inside the direcotory

		// check the no of files returned
		if len(fPaths) == 0 {
			// thorw error
			fmt.Println("> Error: no files found with extention:", e)
			os.Exit(1) // exit the program
		}
	}

	fmt.Printf("> No of file found=%v, in directory=%v, of format=%v\n", len(fPaths), d, e)

	// loop through all the path
	for _, fp := range fPaths {
		tl, f := file.Read(fp)        // get all the list of lines and file
		txt := file.EditLines(&tl, s) // get the edited text
		file.Write(txt, f)            // write the file
		fmt.Println(fp)               // console print the filepath
	}

}
