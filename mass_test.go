package main
import "testing"

func TestCompare500mgWith10g(t *testing.T) {
	mass1 := NewMass(500, MG)
	mass2 := NewMass(10, G)
	result := mass1.Compare(mass2)
	expected := false
	if result != expected {
		t.Errorf("Expected false, but returned true")
	}
}

func TestCompare1000gWith1kg(t *testing.T) {
	mass1 := NewMass(1000, G)
	mass2 := NewMass(1, KG)
	result := mass1.Compare(mass2)
	expected := true
	if result != expected {
		t.Errorf("Expected true but returned false")
	}
}

func TestAdd1kgAnd2000g(t *testing.T) {
	mass1 := NewMass(1, KG)
	mass2 := NewMass(2000, G)
	result := mass1.Add(mass2)
	expected := NewMass(3000, G)
	
	if result != expected {
		t.Errorf("Incorrect addition or unit")
	}
}

func TestSubtract1gWith200mg(t *testing.T) {
	mass1 := NewMass(1, G)
	mass2 := NewMass(200, MG)
	result := mass1.Subtract(mass2)
	expected := NewMass(800, MG)
	if result != expected {
		t.Errorf("Incorrect subtraction or unit")
	}
}

