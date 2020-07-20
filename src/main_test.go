package main

import "testing"

func TestSayHello(t *testing.T) {
	got := sayHello()

	if expected := "hello"; got != expected {
		t.Errorf("expected: %s; got: %s", expected, got)
	}
}
