package main

import (
	"time"
	"os"
	"net/http"
	"fmt"
	"bufio"
	"io"
)

func main() {

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[2:] {
		go fetch(url, ch)
	}

	for range os.Args[2:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {

	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	secs := time.Since(start).Seconds()
	fileName := os.Args[1]
 	file, err := os.Create(fileName)

	if err != nil {
		ch <- fmt.Sprintf("while opening file: %s, %v", fileName, err)
	}

	writer := bufio.NewWriter(file)
	nbytes, err := io.Copy(writer, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while copying %s: %v", url, err)
		return
	}

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}