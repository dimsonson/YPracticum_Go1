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

/* package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type LimitedReader struct {
	reader io.Reader
	//  запоминаем количество считанных байт
	left int
}

func LimitReader(r io.Reader, n int) io.Reader {
	return &LimitedReader{reader: r, left: n}
}

func (r *LimitedReader) Read(p []byte) (int, error) {
	if r.left == 0 {
		return 0, io.EOF
	}
	if r.left < len(p) {
		p = p[0:r.left]
	}
	n, err := r.reader.Read(p)
	r.left -= n
	return n, err
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := LimitReader(r, 4)

	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		log.Fatal(err)
	}
}
*/
