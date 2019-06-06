package serial

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/edenriquez/load-balancer-proxy-go/requestor/flagger"
)

var globalCounter = 0

func makeRequest(c flagger.Commands) {
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
		fmt.Printf(flagger.InfoColor, body)
	}
}

// Requestor should request in parallel to the given url
func Requestor(args ...flagger.Commands) {
	for _, command := range args {
		makeRequest(command)
	}
}
