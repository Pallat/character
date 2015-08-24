package main

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

func LastIChar(s string) counter {
	return counter{s: s}
}
