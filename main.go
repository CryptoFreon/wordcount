package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	text := getInput()
	words := strings.Fields(text)
	fmt.Println(len(words))

}

// getInput return string from StdIn
func getInput() string {
	flag.Parse()
	txt := strings.Join(flag.Args(), "")
	return string(txt)
}
