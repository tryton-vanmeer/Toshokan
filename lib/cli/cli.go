package cli

import (
	"os"
	"strings"
)

func validateArgs() bool {
	return len(os.Args) >= 2
}

func getSearchTerms() string {
	return strings.Join(os.Args[1:], " ")
}

func Run() error {
	if !validateArgs() {
		return nil
	}

	return nil
}
