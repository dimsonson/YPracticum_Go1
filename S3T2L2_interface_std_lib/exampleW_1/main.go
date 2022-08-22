package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

func LimitReader(r io.Reader, n int) io.Reader {
	buf := make([]byte, n)
	r1 := bytes.NewReader(buf)
	_, err := r.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	return r1
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
