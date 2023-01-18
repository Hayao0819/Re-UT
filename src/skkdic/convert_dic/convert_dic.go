package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: convert_dic <path of id.def> <skkdict path>")
		os.Exit(1)
	}

	iddefPath := args[1]
	skkdictPath := args[2]

	err := convert_dic(iddefPath, skkdictPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot convert dic")
		os.Exit(1)
	}
}


func convert_dic(iddefPath, skkdictPath string) (error) {
	// あとで実装
	return nil
}
