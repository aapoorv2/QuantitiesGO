package main
import "testing"

func TestCompare5mWith10cm(t *testing.T) {
	length1 := NewLength(5, M)
	length2 := NewLength(10, CM)
	result := length1.Compare(&length2)
	expected := false
	if result != expected {
		t.Errorf("Expected false, but returned true")
	}
}

func TestCompare10mmWith1cm(t *testing.T) {
	length1 := NewLength(10, MM)
	length2 := NewLength(1, CM)
	result := length1.Compare(&length2)
	expected := true
	if result != expected {
		t.Errorf("Expected true but returned false")
	}
}

func TestAdd1mAnd2cm(t *testing.T) {
	length1 := NewLength(1, M)
	length2 := NewLength(2, CM)
	result := length1.Add(&length2)
	expected := NewLength(102, CM)
	if result != expected {
		t.Errorf("Incorrect addition or unit")
	}
}

func TestSubtract1mWith2cm(t *testing.T) {
	length1 := NewLength(1, M)
	length2 := NewLength(2, CM)
	result := length1.Subtract(&length2)
	expected := NewLength(98, CM)
	if result != expected {
		t.Errorf("Incorrect subtraction or unit")
	}
}

