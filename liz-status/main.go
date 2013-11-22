package main

import (
	"os"
	"os/exec"

	"github.com/seagreen/liz/libs/editor"
)

func main() {

	text := editor.Editor()

	cmd := exec.Command("./liz-tent", text)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
