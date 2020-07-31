package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "Wooo! 	I  LOVE\u00A0   my espaços\t			so much <3"
	fmt.Println(string(sqshSpace([]byte(s))))
}

func sqshSpace(b []byte) []byte {
	j := 0
	for i := 0; i < len(b); {
		r1, rs1 := utf8.DecodeRune(b[j:])
		r2, rs2 := utf8.DecodeRune(b[j+rs1:])
		if unicode.IsSpace(r1) && unicode.IsSpace(r2) {
			b[j] = []byte(" ")[0]
			copy(b[j+1:], b[j+rs1+rs2:])
			j -= rs2 - 1
			i += rs1
			continue
		}
		i += rs1
		j += rs1
	}

	return b[:j]
}
