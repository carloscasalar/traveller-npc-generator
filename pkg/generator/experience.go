package generator

//go:generate enumer -type=Experience -output=experience_auto.go -trimprefix=Experience -transform=snake
type Experience int

const (
	ExperienceRecruit Experience = iota
	ExperienceRookie
	ExperienceIntermediate
	ExperienceRegular
	ExperienceVeteran
	ExperienceElite
)
