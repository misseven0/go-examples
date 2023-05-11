package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/misseven0/httpcache/diskcache"
	"github.com/misseven0/httpcache/httpcache"
)

// https://github.com/misseven0/httpcache
// https://www.golangprograms.com/how-do-you-handle-http-client-caching-in-go.html
func main() {
	// Create a new disk cache with a maximum size of 100MB and a TTL of 1 hour
	cache := diskcache.New("cache")
	// cache.MaxSize(100 * 1024 * 1024)
	// cache.TTL(time.Hour)

	// Create a new HTTP transport with caching enabled
	transport := httpcache.NewTransport(cache)
	client := &http.Client{Transport: transport}

	// Send a GET request to the HTTP server
	req, _ := http.NewRequest("GET", "https://static.oschina.net/new-osc/js/utils/jquery.min.js", nil)
	req.Header.Add("accept-encoding", "gzip, deflate")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
