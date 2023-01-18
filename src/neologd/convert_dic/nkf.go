package main

import (
	"os"
	"os/exec"
	"fmt"
	"strings"
)

func checkNkf() error{
	// get nkf path
	err := exec.Command("which", "nkf").Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "nkf is not installed")
		//os.Exit(1)
		return err
	}
	return nil
}

func kata2hira(s string) (string, error) {
	//checkNkf()
	cmd := exec.Command("nkf", "-w", "--hiragana")
	cmd.Stdin = strings.NewReader(s)
	result, err := cmd.Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return "", err
	}

	return string(result), nil
}

