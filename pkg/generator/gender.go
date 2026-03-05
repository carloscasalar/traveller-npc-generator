package generator

//go:generate go tool enumer -type=Gender -output=gender_auto.go -trimprefix=Gender
type Gender int

const (
	GenderUnspecified Gender = iota
	GenderFemale
	GenderMale
)
