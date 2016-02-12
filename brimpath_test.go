package brimpath

import "testing"

func TestNormalizePath(t *testing.T) {
	out := NormalizePath("a")
	exp := "a"
	if out != exp {
		t.Errorf("%#v != %#v", out, exp)
	}
	out = NormalizePath("a/..")
	exp = "."
	if out != exp {
		t.Errorf("%#v != %#v", out, exp)
	}
	out = NormalizePath("a/../b")
	exp = "b"
	if out != exp {
		t.Errorf("%#v != %#v", out, exp)
	}
	out = NormalizePath("a/../../b")
	exp = "../b"
	if out != exp {
		t.Errorf("%#v != %#v", out, exp)
	}
}
