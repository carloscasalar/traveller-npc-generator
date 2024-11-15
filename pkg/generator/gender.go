package generator

import "github.com/carloscasalar/traveller-npc-generator/internal/name"

//go:generate enumer -type=Gender -output=gender_auto.go -trimprefix=Gender
type Gender int

func (i Gender) toNpcGender() name.Gender {
	switch i {
	case GenderFemale:
		return name.GenderFemale
	case GenderMale:
		return name.GenderMale
	case GenderUnspecified:
		return name.GenderUnspecified
	default:
		return name.Gender(i)
	}
}

const (
	GenderUnspecified Gender = iota
	GenderFemale
	GenderMale
)
