package main

import (
	"ill.fi/neobeam/interp"
	"testing"
	"unicode"
)

func TestCreateRope(t *testing.T) {
	r := interp.NewRope("abc")
	if string(*r.Current) != "a" {
		EShouldBe(t, "Current", string(*r.Current), "a")
	} else if string(r.Source) != "abc" {
		EShouldBe(t, "Source", string(r.Source), "abc")
	}
}

func TestFullConsume(t *testing.T) {
	r := interp.NewRope("123")
	dat := r.Consume(unicode.IsDigit)
	if string(dat) != "123" {
		EShouldBe(t, "Data", string(dat), "123")
	}
}

func TestConsume(t *testing.T) {
	r := interp.NewRope("123abc")
	dat := r.Consume(unicode.IsDigit)
	if string(dat) != "123" {
		EShouldBe(t, "Data", string(dat), "123")
	}
	if string(r.Source) != "abc" {
		EShouldBe(t, "Source after consume", string(r.Source), "")
	}
}
