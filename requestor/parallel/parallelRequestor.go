package parallel

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	infoColor    = "\033[1;34m%s\033[0m"
	noticeColor  = "\033[1;36m%s\033[0m"
	warningColor = "\033[1;33m%s\033[0m"
	errorColor   = "\033[1;31m%s\033[0m"
	debugColor   = "\033[0;36m%s\033[0m"
)

// Commands shold store url´s and request number
type Commands struct {
	URL    string
	Number int64
}

// Requestor should take args and request in parallel
const instructions = `
Usage: go run main.go [options...] <url> , ...

Options:
	-n	number of requests
	-u	indicates the url to request
`

var globalCounter = 0

func makeRequest(c Commands) {
	for index := int64(0); index < c.Number; index++ {
		globalCounter++
		fmt.Printf(
			noticeColor,
			"Starting request number:["+strconv.Itoa(globalCounter)+"] ",
		)
		resp, err := http.Get(c.URL)
		if err != nil {
			fmt.Printf(errorColor, err.Error()+"\n")
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf(infoColor, body)
	}
}

// Requestor should request in parallel to the given url
func Requestor() {
	if len(os.Args[1:]) <= 1 {
		fmt.Println(instructions)
		return
	}
	var args []Commands
	urlObj := Commands{}
	for ind := range os.Args[1:] {
		if os.Args[ind] == "-u" {
			if len(os.Args[ind+1]) < 10 {
				fmt.Printf(errorColor, "[-] Wrong URL format\n")
				os.Exit(0)
			}
			urlObj.URL = os.Args[ind+1] // url
		}
		if os.Args[ind] == "-n" {
			n, _ := strconv.ParseInt(os.Args[ind+1], 10, 0)
			if n < 1 {
				fmt.Printf(errorColor, "[-] Invalid request number\n")
				os.Exit(0)
			}
			urlObj.Number = n // number of requests
		}

		if urlObj.Number > 0 && urlObj.URL != "" {
			args = append(args, urlObj)
			urlObj = Commands{}
		}
	}
	for _, command := range args {
		makeRequest(command)
	}
}
