package main

import (
	"reflect"
	"testing"
)

func TestSplitRepeatChar(t *testing.T) {
	s := "aaaabbbccadd"
	if !reflect.DeepEqual(split(s), []string{"aaaa", "bbb", "cc", "a", "dd"}) {
		t.Error("I need to split string to like this ", []string{"aaaa", "bbb", "cc", "a", "dd"}, " but is was ", split(s))
	}
}

func TestEcho(t *testing.T) {
	p := split("aaaabbbccadd")
	s := echo(p)
	if s != "a3b2c1a0d1" {
		t.Error("I need a3b2c1a0d1 but got ", s)
	}
}
