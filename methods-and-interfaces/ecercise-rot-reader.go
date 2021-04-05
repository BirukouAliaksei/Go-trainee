package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

var ascii_uppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var ascii_lowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var ascii_uppercase_len = len(ascii_uppercase)
var ascii_lowercase_len = len(ascii_lowercase)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	pos := bytes.IndexByte(ascii_uppercase, b)
	if pos != -1 {
		return ascii_uppercase[(pos+13)%ascii_uppercase_len]
	}

	pos = bytes.IndexByte(ascii_lowercase, b)
	if pos != -1 {
		return ascii_lowercase[(pos+13)%ascii_lowercase_len]
	}
	return b
}

func (r rot13Reader) Read(bytes []byte) (n int, err error) {
	n, err = r.r.Read(bytes)
	for i := 0; i < n; i++ {
		bytes[i] = rot13(bytes[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
