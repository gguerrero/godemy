package godog

import (
	"fmt"
	"testing"
)

func BenchmarkToDogYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToDogYears(7)
	}
}

func BenchmarkToDogYears2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ToDogYears2(7)
	}
}

func ExampleToDogYears() {
	fmt.Println(ToDogYears(7))
	// Output:
	// 49
}

func ExampleToDogYears2() {
	fmt.Println(ToDogYears2(7))
	// Output:
	// 49
}

func TestToDogYears(t *testing.T) {
	got := ToDogYears(7)
	expected := 49
	if got != expected {
		t.Error("Got", got, "but expected", expected)
	}
}

func TestToDogYears2(t *testing.T) {
	got := ToDogYears2(7)
	expected := 49
	if got != expected {
		t.Error("Got", got, "but expected", expected)
	}
}
