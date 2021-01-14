package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func domainscan(pattern string, wordfile string, server string) error {
	// Read out words file into array of words
	file, err := os.Open(wordfile)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(1000), // 1s
			}
			return d.DialContext(ctx, "udp", fmt.Sprintf("%s:53", server))
		},
	}

	for scanner.Scan() {
		word := scanner.Text()
		host := strings.Replace(pattern, "%", word, 1)

		addresses, err := resolver.LookupHost(context.Background(), host)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: Not found\n", host)
		} else {
			fmt.Printf("%s\n", host)
		}
	}

	return nil
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
			&cli.StringFlag{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "DNS Server to query",
				Value:   "8.8.4.4",
			},
		},
		Action: func(c *cli.Context) error {
			err := domainscan(c.String("pattern"), c.String("words"), c.String("server"))
			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
