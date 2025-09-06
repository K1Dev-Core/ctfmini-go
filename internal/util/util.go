package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func XOR(data, key []byte) []byte {
	out := make([]byte, len(data))
	for i := range data {
		out[i] = data[i] ^ key[i%len(key)]
	}
	return out
}

func ReadAllStdin() ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, os.Stdin)
	return buf.Bytes(), err
}

func HexDump(b []byte, width int) string {
	if width <= 0 {
		width = 16
	}
	var out bytes.Buffer
	for i := 0; i < len(b); i += width {
		end := i + width
		if end > len(b) {
			end = len(b)
		}
		line := b[i:end]
		fmt.Fprintf(&out, "%08x  ", i)
		for j := 0; j < width; j++ {
			if i+j < len(b) {
				fmt.Fprintf(&out, "%02x ", b[i+j])
			} else {
				out.WriteString("   ")
			}
			if j == 7 {
				out.WriteByte(' ')
			}
		}
		out.WriteString(" |")
		for _, c := range line {
			if c >= 32 && c <= 126 {
				out.WriteByte(c)
			} else {
				out.WriteByte('.')
			}
		}
		out.WriteString("|
")
	}
	return out.String()
}
