package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(target byte) byte {
	const letterCount, rotCount = 26, 13

	switch {
	case target >= 'a' && target <= 'z':
		result := (target - 'a' + rotCount) % letterCount + 'a'
		return result
	case target >= 'A' && target <= 'Z':
		result := (target - 'A' + rotCount) % letterCount + 'A'
		return result
	default:
		return target
	}
}

func (r13reader rot13Reader) Read(target []byte) (bytesRead int, err error) {
	bytesRead, err = r13reader.r.Read(target)

	for i := range target {
		target[i] = rot13(target[i])
	}

	return bytesRead, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
