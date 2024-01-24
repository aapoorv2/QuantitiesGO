package main

type MassUnit float64

const (
	MG MassUnit = 1
	G MassUnit = 1000
	KG MassUnit = 1000000
)

type Mass struct {
	value float64
	unit  MassUnit
}


func NewMass(val float64, unit MassUnit) Mass {
	return Mass{
		value: val,
		unit:  unit,
	}
}

func (m Mass) convertTo(other Mass) float64{
	return m.value * float64(m.unit) / float64(other.unit)
}

func (m Mass) Compare(other Mass) bool {
	mass1 := m.convertTo(other)
	mass2 := other.value
	return mass2 == mass1
}
func (m Mass) Add(other Mass) Mass {
	mass1 := m.convertTo(other)
	mass2 := other.value;
	return Mass{
		value: mass1 + mass2,
		unit: other.unit,
	}
}

func (m Mass) Subtract(other Mass) Mass {
	mass1 := m.convertTo(other)
	mass2 := other.value;
	return Mass{
		value: mass1 - mass2,
		unit: other.unit,
	}
}