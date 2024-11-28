package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("git", "diff", "--name-only", "--diff-filter=U", "--relative")
    var out bytes.Buffer
    cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

    fileList := strings.Split(out.String(), "\n")

    cmd = exec.Command("nvim", fileList...)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
