package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main(){
	for _, url := range os.Args[1:]{
		httpPrefix := "http://"
		completeUrl := ""
		if !strings.HasPrefix(url, httpPrefix) {
			completeUrl = httpPrefix + url
		}

		resp, err := http.Get(completeUrl)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s", resp.Status)
	}
}

