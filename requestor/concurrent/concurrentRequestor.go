package concurrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/edenriquez/load-balancer-proxy-go/requestor/flagger"
)

var globalCounter = 0
var tmp = int64(0)

func makeRequest(c flagger.Commands, ch chan<- string) {
	tmp += c.Number
	go func() {
		for index := int64(0); index < c.Number; index++ {
			client := &http.Client{}
			req, err := http.NewRequest("GET", c.URL, nil)
			req.Header.Set("domain", c.Domain)
			res, err := client.Do(req)
			globalCounter++
			if err != nil {
				fmt.Printf(flagger.InfoColor, "Starting request number:["+strconv.Itoa(globalCounter)+"]"+err.Error()+"\n")
			}
			if res != nil {
				body, _ := ioutil.ReadAll(res.Body)
				ch <- string(body)
				res.Body.Close()
				fmt.Printf(flagger.InfoColor, "Starting request number:["+strconv.Itoa(globalCounter)+"]\n")
			}
			if tmp-1 == int64(globalCounter) || tmp == int64(globalCounter) {
				fmt.Println("closing ")
				close(ch)
			}
		}
	}()

}

// Requestor should execute concurrent requests
func Requestor(args ...flagger.Commands) {
	ch := make(chan string)
	go func() {
		for _, command := range args {
			makeRequest(command, ch)
		}
	}()

	for c := range ch {
		fmt.Println(c)
	}

}
