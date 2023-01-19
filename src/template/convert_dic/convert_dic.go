package main

import (
	"os"
	"fmt"
	//"strings"
	//common "github.com/Hayao0819/Re-UT/common"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: convert_dic <path of id.def> <dict path>")
		os.Exit(1)
	}
	//iddefPath := args[1]
	//dictPath := args[2]
}

