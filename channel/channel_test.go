package channel

import (
	"testing"
)

func Test(t *testing.T) {
	c := make(chan Foo, 10)
	c <- Baz{"Foo"}
	if DoIt(c) != "burp" {
		t.Error("Expected burp")
	}

}
