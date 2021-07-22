package cli

import (
	"fmt"
	"os"
	"strings"
)

func validateArgs() bool {
	return len(os.Args) >= 2
}

func Run() {
	if !validateArgs() {
		fmt.Println("Not enough arguments")
	} else {
		fmt.Printf("SEARCH: %s", strings.Join(os.Args[1:], " "))
	}
}
