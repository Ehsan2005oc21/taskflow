package main

import (
	"fmt"
	"os"
)

// version is printed at startup.
const version = "0.1.0"
func main() {
	fmt.Println("TaskFlow is runing, version", version)
	if len(os.Args) > 1 {
		fmt.Println("args:", os.Args[1:])
	}
	os.Exit(run())
}

func run() int {
	tasks := []string{"write README", "add tests"}
	for i, t := range tasks {
		fmt.Printf("%d. %s\n", i+1, t)
	}
	return 0
}
