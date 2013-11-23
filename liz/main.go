package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	arg := os.Args[1]
	if arg == "status" {

		cmd := exec.Command("./liz-status")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

	} else {
		fmt.Println(`Available commands: "status"`)
	}
}
