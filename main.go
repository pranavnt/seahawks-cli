package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You have to provide a command :)")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "score":
		fmt.Println("The seahawks score is...")
	default:
		fmt.Println("command not recognized")
		os.Exit(1)
	}
}
