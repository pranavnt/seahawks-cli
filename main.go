package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("must provide command")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "score":
		fmt.Println(os.Args[1])
	default:
		fmt.Println("command not recognized")
		os.Exit(1)
	}
}
