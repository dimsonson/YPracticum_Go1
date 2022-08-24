package main

import "fmt"

func main1() {
	var ret string
	fmt.Scan(&ret)
	if ret == `q` {
		fmt.Print("testtt")
	}
}
