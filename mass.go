package main

type Mass struct {
	value float64
	unit  string
	unitToFactor map[string]float64
}


func NewMass(val float64, unit string) Mass {
	unitToFactor := map[string]float64{
		"MG": 1,
		"G": 1000,
		"KG": 1_000_000,
	}
	return Mass{
		value: val,
		unit:  unit,
		unitToFactor: unitToFactor,
	}
}

func (m *Mass) convertTo(other *Mass) float64{
	return m.value * m.unitToFactor[m.unit] / m.unitToFactor[other.unit]
}

func (m *Mass) Compare(other *Mass) bool {
	mass1 := m.convertTo(other)
	mass2 := other.value
	return mass2 == mass1
}
func (m *Mass) Add(other *Mass) float64 {
	mass1 := m.convertTo(other)
	mass2 := other.value;
	return mass1 + mass2
}

func (m *Mass) Subtract(other *Mass) float64 {
	mass1 := m.convertTo(other)
	mass2 := other.value;
	return mass1 - mass2
}