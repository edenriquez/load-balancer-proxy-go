// concurrent.go
package concurrent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/edenriquez/load-balancer-proxy-go/requestor/flagger"
)

func MakeRequest(url string, ch chan<- string) {

	resp, _ := http.Get(url)

	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}

func Requestor(args ...flagger.Commands) {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go MakeRequest(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}
