package main

import (
	"fmt"
	"strconv"
)

type counter struct {
	s string
}

func split(s string) []string {
	p := []string{}
	var previous rune
	for i, v := range s {
		if i == 0 {
			previous = v
			continue
		}
		if v != previous {
			p = append(p, s[:i])
			s = s[i:]
			p = append(p, split(s)...)
			break
		}

		if i == len(s)-1 {
			p = append(p, s)
		}
	}
	return p
}

func lastIndex(s string) string {
	return string(s[0]) + strconv.Itoa(len(s)-1)
}

func echo(p []string) string {
	s := ""
	for _, v := range p {
		s += lastIndex(v)
	}
	return s
}

func main() {
	fmt.Println(echo(split("aaaabbbccaddeeeeedfedffffff")))
}
