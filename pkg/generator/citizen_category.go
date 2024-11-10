package generator

//go:generate enumer -type=CitizenCategory -output=citizen_category_auto.go -trimprefix=Category -transform=snake

type CitizenCategory int

const (
	CategoryBelowAverage CitizenCategory = iota
	CategoryAverage
	CategoryAboveAverage
	CategoryExceptional
)
