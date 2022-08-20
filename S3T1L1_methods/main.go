package main

import (
	"time"
	"fmt"
)

import

func main() {
    sw := Stopwatch{}
    sw.Start()

    time.Sleep(1 * time.Second)
    sw.SaveSplit()

    time.Sleep(500 * time.Millisecond)
    sw.SaveSplit()

    time.Sleep(300 * time.Millisecond)
    sw.SaveSplit()

    fmt.Println(sw.GetResults())
} 