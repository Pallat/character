package main

type counter struct {
	s string
}

func (c counter) piece() []string {
	return []string{}
}

func LastIChar(s string) counter {
	return counter{s: s}
}
