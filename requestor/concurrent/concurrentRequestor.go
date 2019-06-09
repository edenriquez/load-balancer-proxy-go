package concurrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/edenriquez/load-balancer-proxy-go/requestor/flagger"
)

var globalCounter = 0

func makeRequest(c flagger.Commands, ch chan<- string) {
	for index := int64(0); index < c.Number; index++ {
		globalCounter++
		fmt.Printf(
			flagger.NoticeColor,
			"Starting request number:["+strconv.Itoa(globalCounter)+"] ",
		)
		resp, err := http.Get(c.URL)
		if err != nil {
			fmt.Printf(flagger.ErrorColor, err.Error()+"\n")
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)

		ch <- string(body)

	}
}

// Requestor should execute concurrent requests
func Requestor(args ...flagger.Commands) {
	ch := make(chan string)
	for _, command := range args {
		go makeRequest(command, ch)
	}

	for j := 0; j < len(args); j++ {
		fmt.Println(<-ch)
	}
}
