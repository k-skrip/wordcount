package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var str []string
	var args string
	if len(os.Args) > 1 {
		for i := range os.Args {
			if i == 0 {
				continue
			}
			args = args + " " + os.Args[i]
		}
	}
	str = strings.Fields(args)
	fmt.Println(len(str))
}
