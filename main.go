package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatalln("problem reading os.stdin")
	}

    // don't work with anything that is not piped
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("bye!")
		os.Exit(0)
	}

	cons, _ := io.ReadAll(os.Stdin)
    conNewLine := strings.Split(string(cons), "\n")
    writer := bufio.NewWriter(os.Stdout)
    var buffer bytes.Buffer

    for idx, line := range conNewLine {
        if idx == len(conNewLine)-2 {
            break
        }
        file := strings.Split(line[26:], " ")[0]
        buffer.WriteString(file + " ")
    }

    writer.WriteString(buffer.String())
    writer.Flush()
}
