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
		go func() {

			resp, err := http.Get(c.URL)
			if err != nil {
				fmt.Printf(flagger.ErrorColor, err.Error()+"\n")
			}
			globalCounter++
			fmt.Printf(
				flagger.NoticeColor,
				"Starting request number:["+strconv.Itoa(globalCounter)+"] ",
			)
			if resp != nil {
				body, _ := ioutil.ReadAll(resp.Body)
				ch <- string(body)
				resp.Body.Close()
			}

		}()
	}
}

// Requestor should execute concurrent requests
func Requestor(args ...flagger.Commands) {
	ch := make(chan string)
	go func() {
		for _, command := range args {
			makeRequest(command, ch)
		}
		// close(ch)
	}()

	for c := range ch {
		fmt.Println(c)
	}

}
