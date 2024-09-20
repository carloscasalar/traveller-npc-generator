package npc

import "fmt"

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

type levelDistribution struct {
	numberOfSkillsLevel0 int
	numberOfSkillsLevel1 int
	numberOfSkillsLevel2 int
	numberOfSkillsLevel3 int
}

func (d levelDistribution) DistributeLevels(skills []string) []string {
	level3Skills, skills := pop(skills, d.numberOfSkillsLevel3)
	level2Skills, skills := pop(skills, d.numberOfSkillsLevel2)
	level1Skills, skills := pop(skills, d.numberOfSkillsLevel1)
	level0Skills, skills := pop(skills, d.numberOfSkillsLevel0)

	skillsWithLevels := make([]string, 0, len(level0Skills)+len(level1Skills)+len(level2Skills)+len(level3Skills))
	skillsWithLevels = append(skillsWithLevels, appendLevel(level3Skills, 3)...)
	skillsWithLevels = append(skillsWithLevels, appendLevel(level2Skills, 2)...)
	skillsWithLevels = append(skillsWithLevels, appendLevel(level1Skills, 1)...)
	skillsWithLevels = append(skillsWithLevels, appendLevel(level0Skills, 0)...)
	return skillsWithLevels
}

var skillLevelDistributionByExperience = map[Experience]levelDistribution{
	ExperienceRecruit: {
		numberOfSkillsLevel0: 4,
	},
	ExperienceRookie: {
		numberOfSkillsLevel0: 4,
		numberOfSkillsLevel1: 2,
	},
	ExperienceIntermediate: {
		numberOfSkillsLevel0: 4,
		numberOfSkillsLevel1: 2,
		numberOfSkillsLevel2: 1,
	},
	ExperienceRegular: {
		numberOfSkillsLevel0: 5,
		numberOfSkillsLevel1: 2,
		numberOfSkillsLevel2: 2,
	},
	ExperienceVeteran: {
		numberOfSkillsLevel0: 5,
		numberOfSkillsLevel1: 2,
		numberOfSkillsLevel2: 3,
	},
	ExperienceElite: {
		numberOfSkillsLevel0: 6,
		numberOfSkillsLevel1: 3,
		numberOfSkillsLevel2: 2,
		numberOfSkillsLevel3: 1,
	},
}

func appendLevel(skills []string, level int) []string {
	skillsWithLevels := make([]string, len(skills))
	for i, skill := range skills {
		skillsWithLevels[i] = fmt.Sprintf("%s-%d", skill, level)
	}
	return skillsWithLevels
}

func pop(skills []string, numberOfSkillsToPop int) (resultingSkills []string, remaining []string) {
	if len(skills) == 0 {
		return []string{}, []string{}
	}
	if numberOfSkillsToPop > len(skills) {
		numberOfSkillsToPop = len(skills)
	}
	resultingSkills = skills[:numberOfSkillsToPop]
	remaining = skills[numberOfSkillsToPop:]
	return
}
