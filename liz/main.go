package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/seagreen/liz/editor"
)

func main() {
	arg := os.Args[1]
	if arg == "status" {

		text := editor.Editor()

		cmd := exec.Command("./liz-tent", text)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

	} else {
		fmt.Println("Available commands: \"status\"")
	}
}
