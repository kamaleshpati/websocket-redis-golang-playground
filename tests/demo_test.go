package tests

import (
	"strconv"
	"testing"
)

func Fooer(input int) string {
	isfoo := (input % 3) == 0
	if isfoo {
		return "Foo"
	}
	return strconv.Itoa(input)
}

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
	}
}
