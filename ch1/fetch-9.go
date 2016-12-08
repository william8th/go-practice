package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

func main() {

	const httpPrefix = "http://"

	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "fetch status: %s\n", resp.Status)

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
