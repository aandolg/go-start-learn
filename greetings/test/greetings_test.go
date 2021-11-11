package test

import (
	"regexp"
	"testing"

	"good-work.in/greetings"
)

func TestHelloName(t *testing.T) {
	name := "World"
	want := regexp.MustCompile(name)
	msg, err := greetings.Hello(name)

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello(%q) = %q, %v, want match for %#q, nil`, name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := greetings.Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
