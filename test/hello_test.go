package greet

import "testing"

// go test -cpuprofile=cprof, then pprof cprof, then web
func TestGreet(t *testing.T) {
	greeting := "Hello"
	if greeting != "Hello" {
		t.Error("Test failed")
	}
}
