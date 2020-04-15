package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yumyum-pi/txtBP/internal"
)

var e string        // file extention to search for
var d string = "./" // directory to find the file in
var s string = "-"  // seperator of the line

// create the root cobra command
var rootCmd = &cobra.Command{
	Use: "txtBP", // name fo the root cli command
	// about the commnad
	Short: "txtBP is a automation tool to number the bullet points in a text file",
	Long:  `txtBP is a automation tool to number the bullet points in a text file. Large`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			d = args[0]
		}

		internal.Start(d, e, s)
	},
}

// Execute runs the main cli functions
func Execute() {
	// check for error
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err) // print the errror
		os.Exit(1)       // exit the program
	}
}

func init() {
	// setting the flags
	// getting the
	rootCmd.PersistentFlags().StringVarP(&e, "extension", "e", ".txt", "File format of the files to read. Default=.txt")
	rootCmd.PersistentFlags().StringVarP(&s, "sep", "s", "-", "Seperator of the line. Default='-'")
}
