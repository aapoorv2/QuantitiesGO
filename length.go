package main
type LengthUnit float64

const (
	MM LengthUnit = 1
	CM LengthUnit = 10
	M  LengthUnit = 1000
	KM  LengthUnit = 1000000
)


type Length struct {
	value float64
	unit  LengthUnit
}

func NewLength(val float64, unit LengthUnit) Length {
	return Length{
		value: val,
		unit:  unit,
	}
}

func (l Length) convertTo(other LengthUnit) float64{
	return l.value * float64(l.unit) / float64(other)
}

func (l Length) Compare(other Length) bool {
	len1 := l.convertTo(other.unit)
	len2 := other.value
	return len1 == len2
}
func (l Length) Add(other Length) Length {
	len1 := l.convertTo(other.unit)
	len2 := other.value;
	return Length {
		value: len1 + len2,
		unit: other.unit,
	}
}

func (l Length) Subtract(other Length) Length {
	len1 := l.convertTo(other.unit)
	len2 := other.value;
	return Length {
		value: len1 - len2,
		unit: other.unit,
	}
}