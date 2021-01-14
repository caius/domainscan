package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
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
	app := &cli.App{
		Name:  "domainscan",
		Usage: "Scan for registered domain names with apex DNS record",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "words",
				Aliases:  []string{"w"},
				Usage:    "word file for pattern replacement, one replacement per line",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "pattern",
				Aliases:  []string{"p"},
				Usage:    "Pattern to inject words into",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			pattern := c.String("pattern")

			// Read out words file into array of words
			file, err := os.Open(c.String("words"))
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				word := scanner.Text()
				candidate := strings.Replace(pattern, "%", word, 1)
				resolveAndOutput(candidate)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
