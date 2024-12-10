package generator

import "github.com/carloscasalar/traveller-npc-generator/internal/npc"

//go:generate enumer -type=CitizenCategory -output=citizen_category_auto.go -trimprefix=CitizenCategory -transform=snake

type CitizenCategory int

func (i CitizenCategory) toNpcCitizenCategory() npc.CitizenCategory {
	switch i {
	case CitizenCategoryBelowAverage:
		return npc.CategoryBelowAverage
	case CitizenCategoryAverage:
		return npc.CategoryAverage
	case CitizenCategoryAboveAverage:
		return npc.CategoryAboveAverage
	case CitizenCategoryExceptional:
		return npc.CategoryExceptional
	default:
		return npc.CitizenCategory(i)
	}
}

const (
	CitizenCategoryBelowAverage CitizenCategory = iota
	CitizenCategoryAverage
	CitizenCategoryAboveAverage
	CitizenCategoryExceptional
)
