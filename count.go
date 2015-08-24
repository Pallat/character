package main

import (
	"fmt"
	"strconv"
)

type counter struct {
	s string
}

func piece(s string) []string {
	p := []string{}
	var r rune
	for i, v := range s {
		if i == 0 {
			r = v
			continue
		}
		if v != r {
			p = append(p, s[:i])
			s = s[i:]
			p = append(p, piece(s)...)
			break
		}

		if i == len(s)-1 {
			p = append(p, s)
		}
	}
	return p
}

func echo(p []string) string {
	s := ""
	for _, v := range p {
		s += string(v[0]) + strconv.Itoa(len(v)-1)
	}
	return s
}

func main() {
	fmt.Println(echo(piece("aaaabbbccadd")))
}
