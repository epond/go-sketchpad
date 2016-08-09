package slice

import "testing"

func TestConstruction(t *testing.T) {
	slice := Construction()
	if len(slice) != 10 {
		t.Error("Expected length to be 10")
	}
	if slice[0] != 3 || slice[1] != 3 || slice[2] != 3 {
		t.Error("Expected all values to be 3")
	}
}

func TestMap(t *testing.T) {
	slice := Map()
	if slice[0] != 2 || slice[1] != 4 || slice[2] != 6 {
		t.Error("Expected map to double values")
	}
}
