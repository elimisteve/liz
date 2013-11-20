package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func Editor() string {
	temp, _ := ioutil.TempFile("", "test")

	cmd := exec.Command("vim", temp.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	bytes, _ := ioutil.ReadFile(temp.Name())

	os.Remove(temp.Name())

	return string(bytes)
}
