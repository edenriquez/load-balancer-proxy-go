package flagger

import (
	"fmt"
	"os"
	"strconv"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

// Commands shold store url´s and request number
type Commands struct {
	URL    string
	Number int64
	Domain string
}

// Requestor should take args and request in parallel
const instructions = `
Usage: go run main.go [options...] <url> , ...
example: go run main.go -u http://localhost -n 2 -h omega -u http://example.com -n 1 ...
Options:
	-h  domain to request [omega, alpha, beta]
	-n	number of requests
	-u	indicates the url to request
`

// Flagger should return the collected data from given arguments
func Flagger() []Commands {
	if len(os.Args[1:]) <= 1 {
		fmt.Println(instructions)
		os.Exit(0)
	}
	var args []Commands
	urlObj := Commands{}
	for ind := range os.Args[1:] {
		if os.Args[ind] == "-u" {
			if len(os.Args[ind+1]) < 10 {
				fmt.Printf(ErrorColor, "[-] Wrong URL format\n")
				os.Exit(0)
			}
			urlObj.URL = os.Args[ind+1] // url
		}
		if os.Args[ind] == "-n" {
			n, _ := strconv.ParseInt(os.Args[ind+1], 10, 0)
			if n < 1 {
				fmt.Printf(ErrorColor, "[-] Invalid request number\n")
				os.Exit(0)
			}
			urlObj.Number = n
		}

		if os.Args[ind] == "-h" {
			n := os.Args[ind+1]
			if len(n) < 1 {
				fmt.Printf(ErrorColor, "[-] Invalid request headers\n")
				os.Exit(0)
			}
			urlObj.Domain = n
		}
		if urlObj.Number > 0 && urlObj.URL != "" && urlObj.Domain != "" {
			args = append(args, urlObj)
			urlObj = Commands{}
		}
	}
	return args
}
