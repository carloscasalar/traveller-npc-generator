package npc

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

func (d levelDistribution) DistributeLevels(skills []skill) []string {
	level3Skills, skills := pop(skills, d.numberOfSkillsLevel3)
	level2Skills, skills := pop(skills, d.numberOfSkillsLevel2)
	level1Skills, skills := pop(skills, d.numberOfSkillsLevel1)

	skills = removeSpecializations(skills)

	skills = make(maxLevelBySkillName).
		SetSkillsLevel(level3Skills, 3).
		SetSkillsLevel(level2Skills, 2).
		SetSkillsLevel(level1Skills, 1).
		RemoveAnySkillWithLevelFrom(skills)

	level0Skills, skills := pop(skills, d.numberOfSkillsLevel0)

	skillsWithLevels := make([]string, 0, len(level0Skills)+len(level1Skills)+len(level2Skills)+len(level3Skills))
	skillsWithLevels = append(skillsWithLevels, appendLevel(level3Skills, 3)...)
	skillsWithLevels = append(skillsWithLevels, appendLevel(level2Skills, 2)...)
	skillsWithLevels = append(skillsWithLevels, appendLevel(level1Skills, 1)...)
	skillsWithLevels = append(skillsWithLevels, appendLevel(level0Skills, 0)...)
	return skillsWithLevels
}

func removeSpecializations(skills []skill) []skill {
	existing := map[string]bool{}
	uniqueSkills := make([]skill, 0, len(skills))
	for _, s := range skills {
		if _, found := existing[s.Name]; found {
			continue
		}
		existing[s.Name] = true
		s.Specialization = noSpecialization
		uniqueSkills = append(uniqueSkills, s)
	}
	return uniqueSkills
}

type maxLevelBySkillName map[string]int

func (m maxLevelBySkillName) SetSkillsLevel(skills []skill, level int) maxLevelBySkillName {
	for _, skill := range skills {
		if level > m[skill.Name] {
			m[skill.Name] = level
		}
	}
	return m
}

func (m maxLevelBySkillName) RemoveAnySkillWithLevelFrom(skills []skill) []skill {
	filteredSkills := make([]skill, 0, len(skills))
	for _, s := range skills {
		if _, found := m[s.Name]; found {
			continue
		}
		filteredSkills = append(filteredSkills, s)
	}
	return filteredSkills
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

func appendLevel(skills []skill, level int) []string {
	skillsWithLevels := make([]string, len(skills))
	for i, s := range skills {
		skillsWithLevels[i] = s.StringLevel(level)
	}
	return skillsWithLevels
}

func pop(skills []skill, numberOfSkillsToPop int) (resultingSkills []skill, remaining []skill) {
	if len(skills) == 0 {
		return []skill{}, []skill{}
	}
	if numberOfSkillsToPop > len(skills) {
		numberOfSkillsToPop = len(skills)
	}
	resultingSkills = skills[:numberOfSkillsToPop]
	remaining = skills[numberOfSkillsToPop:]
	return
}
