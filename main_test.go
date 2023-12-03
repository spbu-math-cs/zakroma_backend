package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello Golang"

	got := "Hello Golang"

	if want != got {
		t.Fatalf("want %s, got %s\n", want, got)
	}
}
