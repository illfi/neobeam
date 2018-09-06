package main

import (
	"fmt"
	"neobeam/interp"
	"testing"
	"unicode"
)

func ShouldBe(name, src, check string) string {
	return fmt.Sprintf("%s failed, is %q, should be %q", name, src, check)
}

func EShouldBe(t *testing.T, name, src, check string) {
	t.Errorf(ShouldBe(name, src, check))
}

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
