package main

import (
	"bufio"
	"fmt"
	"os"

	"unescape"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	for sc.Scan() {
		fmt.Println(unescape.Unescape(sc.Text()))
	}
}
