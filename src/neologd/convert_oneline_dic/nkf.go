package main

import (
	"os"
	"os/exec"
	"fmt"
	"strings"
)

func checkNkf(){
	// get nkf path
	err := exec.Command("which", "nkf").Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "nkf is not installed")
		os.Exit(1)
	}
}

func kata2hira(s string) string {
	checkNkf()
	cmd := exec.Command("nkf", "-w", "--hiragana")
	cmd.Stdin = strings.NewReader(s)
	result, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(result))
		os.Exit(1)
	}
	return string(result)
}

