package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

const (
	copyCmd  string = "pbcopy"
	pasteCmd string = "pbpaste"
)

func main() {
	var temp string
	for {
		time.Sleep(1 * time.Second)
		out, err := exec.Command(pasteCmd).Output()
		if err != nil {
			log.Fatalln(err)
		}

		text := string(out)
		if text != "" && text != temp {
			fmt.Println(text)
			temp = text
		}
	}
}
