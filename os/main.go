package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Command line arguments: ")
	for i, v := range os.Args {
		fmt.Println(fmt.Sprintf("position: %d - value: %s", i, v))
	}

	fmt.Println("----------")

	fmt.Println("Additional arguments: " + strings.Join(os.Args[1:], ", "))

}
