package main

import (
	"reflect"
	"testing"
)

func TestSplitRepeatChar(t *testing.T) {
	c := LastIChar("aaaabbbccadd")
	if !reflect.DeepEqual(c.piece(), []string{"aaaa", "bbb", "cc", "a", "dd"}) {
		t.Error("I need to split string to like this ", []string{"aaaa", "bbb", "cc", "a", "dd"})
	}
}
