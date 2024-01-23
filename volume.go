package main

type VolumeUnit float64

const (
	CUBIC_MM VolumeUnit = 1
	CUBIC_CM VolumeUnit = 1000
	CUBIC_M VolumeUnit = 1_000_000_000
)

type Volume struct {
	value float64
	unit  VolumeUnit
}

func NewVolume(val float64, unit VolumeUnit) Volume {
	return Volume{
		value: val,
		unit:  unit,
	}
}

func (v *Volume) convertTo(other *Volume) float64{
	return v.value * float64(v.unit) / float64(other.unit)
}

func (v *Volume) Compare(other *Volume) bool {
	vol1 := v.convertTo(other)
	vol2 := other.value
	return vol1 == vol2
}
func (v *Volume) Add(other *Volume) Volume {
	vol1 := v.convertTo(other)
	vol2 := other.value;
	return Volume {
		value: vol1 + vol2,
		unit: other.unit,
	}
}

func (v *Volume) Subtract(other *Volume) Volume {
	vol1 := v.convertTo(other)
	vol2 := other.value;
	return Volume {
		value: vol1 - vol2,
		unit: other.unit,
	}
}