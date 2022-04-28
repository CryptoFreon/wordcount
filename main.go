package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text := getInput()
	words := strings.Fields(text)
	fmt.Println(len(words))

}

// getInput return string from StdIn
func getInput() string {
	scanner := bufio.NewReader(os.Stdin)
	txt, _, _ := scanner.ReadLine()
	return string(txt)
}
