package slice

import "testing"

func TestConstruction(t *testing.T) {
	s := Construction()
	if len(s) != 10 {
		t.Error("Expected length to be 10")
	}
}

func TestMap(t *testing.T) {
	slice := Map()
	if slice[0] != 2 || slice[1] != 4 || slice[2] != 6 {
		t.Error("Expected map to double values")
	}
}
