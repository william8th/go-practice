package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	inefficientStart := time.Now()
	s, sep := "", " "
	for _, v := range os.Args[1:] {
		s += v + sep
	}
	inefficientEnd := time.Since(inefficientStart).Seconds()

	start := time.Now()
	joinS := strings.Join(os.Args[1:], " ")
	end := time.Since(start).Seconds()

	fmt.Printf("Inefficient version took: %fs with value: %s\n", inefficientEnd, s)
	fmt.Printf("Join version took: %fs with value: %s\n", end, joinS)
}
