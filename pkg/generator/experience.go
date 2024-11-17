package generator

import "github.com/carloscasalar/traveller-npc-generator/internal/npc"

//go:generate enumer -type=Experience -output=experience_auto.go -trimprefix=Experience -transform=snake
type Experience int

func (i Experience) toNpcExperience() npc.Experience {
	switch i {
	case ExperienceRecruit:
		return npc.ExperienceRecruit
	case ExperienceRookie:
		return npc.ExperienceRookie
	case ExperienceIntermediate:
		return npc.ExperienceIntermediate
	case ExperienceRegular:
		return npc.ExperienceRegular
	case ExperienceVeteran:
		return npc.ExperienceVeteran
	case ExperienceElite:
		return npc.ExperienceElite
	default:
		return npc.Experience(i)
	}
}

const (
	ExperienceRecruit Experience = iota
	ExperienceRookie
	ExperienceIntermediate
	ExperienceRegular
	ExperienceVeteran
	ExperienceElite
)
