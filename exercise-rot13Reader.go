package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := 0; i < n; i++ {
		if b[i] <= 90 && b[i] >= 65 {
			b[i] = 64 + ((b[i] + 13 - 64) % 26)
		} else if b[i] <= 122 && b[i] >= 97 {
			b[i] = 96 + ((b[i] + 13 - 96) % 26)
		}
	}

	return n, err
}

func main() {
	a := "Lbh penpxrq gur pbqr!"
	s := strings.NewReader(a)
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
