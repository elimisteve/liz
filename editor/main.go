package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	temp, _ := ioutil.TempFile("", "test")

	cmd := exec.Command("vim", temp.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

	fmt.Println(ioutil.ReadFile(temp.Name()))

	os.Remove(temp.Name())
}
