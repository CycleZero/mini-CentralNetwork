package main

import (
	"fmt"

	"./temp"
)

func main() {
	ch := make(chan string)
	var d string

	go temp.Testchan(ch)
	for {
		_, _ = fmt.Scan(&d)
		ch <- d
		if d == "0" {
			break
		}
	}

}
