package main
import "testing"

func TestCompare5cubicmmWith10cubicm(t *testing.T) {
	volume1 := NewVolume(5, CUBIC_MM)
	volume2 := NewVolume(10, CUBIC_CM)
	result := volume1.Compare(volume2)
	expected := false
	if result != expected {
		t.Errorf("Expected false, but returned true")
	}
}

func TestCompare1000cubicmmWith1cubiccm(t *testing.T) {
	volume1 := NewVolume(1000, CUBIC_MM)
	volume2 := NewVolume(1, CUBIC_CM)
	result := volume1.Compare(volume2)
	expected := true
	if result != expected {
		t.Errorf("Expected true, but returned false")
	}
}

func TestAdd1cubicmAnd2cubiccm(t *testing.T) {
	volume1 := NewVolume(1, CUBIC_M)
	volume2 := NewVolume(2, CUBIC_CM)
	result := volume1.Add(volume2)
	expected := NewVolume(1000002, CUBIC_CM)
	
	if result != expected {
		t.Errorf("Incorrect addition or unit")
	}
}

func TestSubtract1cubicmWith2cubiccm(t *testing.T) {
	volume1 := NewVolume(1, CUBIC_M)
	volume2 := NewVolume(2, CUBIC_CM)
	result := volume1.Subtract(volume2)
	expected := NewVolume(999998, CUBIC_CM)
	if result != expected {
		t.Errorf("Incorrect subtraction or unit")
	}
}

