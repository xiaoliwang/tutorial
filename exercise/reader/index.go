package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func (rot *rot13Reader) Read(p []byte) (int, error) {
    n, err := rot.r.Read(p)
    for i := 0; i < len(p); i++ {
        if (p[i] >= 'A' && p[i] < 'N') || (p[i] >='a' && p[i] < 'n') {
            p[i] += 13
        } else if (p[i] > 'M' && p[i] <= 'Z') || (p[i] > 'm' && p[i] <= 'z'){
            p[i] -= 13
        }
    }
    return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader
    io.Copy(os.Stdout, &r)
}
