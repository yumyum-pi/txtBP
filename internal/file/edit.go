package file

import (
	"fmt"
	"strings"
)

// EditLines edits and return the a combine string
func EditLines(txtlines *[]string, s string) (text string) {
	// loop through all the lines
	for i, l := range *txtlines {
		var ln string // number of the line
		// add 0 to number smaller than 9
		if i < 9 {
			ln = fmt.Sprintf("0%v%v", i+1, s)
		} else {
			ln = fmt.Sprintf("%v%v", i+1, s)
		}
		// add the line to the main text
		text += (strings.Replace(l, s, ln, 1) + "\n")
	}

	return
}
