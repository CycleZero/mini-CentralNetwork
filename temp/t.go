package temp

import "fmt"

func Testchan(ch chan string) {
	var data string
	for {
		data = <-ch
		if data != "0" {
			fmt.Println("received:" + fmt.Sprint(data))
		} else {
			fmt.Println("exit")
			break
		}

	}

}
