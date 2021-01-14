package main

import (
	"fmt"
	// flag "github.com/spf13/pflag"
	"net"
	"os"
	"strings"
)

func resolveAndOutput(host string) {
	_, err := net.LookupHost(host)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: Not found\n", host)
		return
	}

	fmt.Printf("%s\n", host)
}

func main() {
	pattern := "caius%%.com"

	replacements := []string{"theory", "3", "2"}

	for _, word := range replacements {
		candidate := strings.Replace(pattern, "%%", word, 1)

		resolveAndOutput(candidate)
	}
}
