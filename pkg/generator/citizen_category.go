package generator

import "github.com/carloscasalar/traveller-npc-generator/internal/npc"

//go:generate enumer -type=CitizenCategory -output=citizen_category_auto.go -trimprefix=Category -transform=snake

type CitizenCategory int

func (i CitizenCategory) toNpcCitizenCategory() npc.CitizenCategory {
	switch i {
	case CategoryBelowAverage:
		return npc.CategoryBelowAverage
	case CategoryAverage:
		return npc.CategoryAverage
	case CategoryAboveAverage:
		return npc.CategoryAboveAverage
	case CategoryExceptional:
		return npc.CategoryExceptional
	default:
		return npc.CitizenCategory(i)
	}
}

const (
	CategoryBelowAverage CitizenCategory = iota
	CategoryAverage
	CategoryAboveAverage
	CategoryExceptional
)
