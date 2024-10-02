package npc

//go:generate enumer -type=CitizenCategory -output=citizen_category_auto.go -trimprefix=Category -transform=snake
type CitizenCategory int

const (
	CategoryBelowAverage CitizenCategory = iota
	CategoryAverage
	CategoryAboveAverage
	CategoryExceptional
)

var characteristicArrayByCitizenCategory = map[CitizenCategory][]int{
	CategoryBelowAverage: {8, 7, 6, 6, 5, 4},
	CategoryAverage:      {9, 8, 7, 7, 6, 5},
	CategoryAboveAverage: {10, 9, 8, 8, 7, 6},
	CategoryExceptional:  {11, 10, 9, 9, 8, 7},
}
