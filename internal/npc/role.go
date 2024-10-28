package npc

import "math/rand"

//go:generate enumer -type=Role -output=role_auto.go -trimprefix=Role -transform=snake
type Role int

const (
	RolePilot Role = iota
	RoleNavigator
	RoleEngineer
	RoleSteward
	RoleMedic
	RoleMarine
	RoleGunner
	RoleScout
	RoleTechnician
	RoleLeader
	RoleDiplomat
	RoleEntertainer
	RoleTrader
	RoleThug
)

func (r Role) Skills(experience Experience) []string {
	if distribution, found := skillLevelDistributionByExperience[experience]; found {
		return distribution.DistributeLevels(r.relevantSkills())
	}
	return []string{}
}

func (r Role) RandomCharacteristic(category CitizenCategory) map[Characteristic]int {
	characteristicsArray := characteristicArrayByCitizenCategory[category]
	characteristics := make(map[Characteristic]int)
	preferences := characteristicPreferenceByRole[r]

	highestValues, remaining := popAndShuffleTwoHighestValues(characteristicsArray)
	for i, characteristic := range preferences.High {
		characteristics[characteristic] = highestValues[i]
	}

	mediumValues, remaining := popAndShuffleTwoHighestValues(remaining)
	for i, characteristic := range preferences.Medium {
		characteristics[characteristic] = mediumValues[i]
	}

	lowestValues, _ := popAndShuffleTwoHighestValues(remaining)
	for i, characteristic := range preferences.Low {
		characteristics[characteristic] = lowestValues[i]
	}

	return characteristics
}

func popAndShuffleTwoHighestValues(values []int) (highest []int, remaining []int) {
	highest = []int{values[0], values[1]}
	rand.Shuffle(len(highest), func(i, j int) {
		highest[i], highest[j] = highest[j], highest[i]
	})
	return highest, values[2:]
}

func (r Role) relevantSkills() []skill {
	if skills, found := skillsByRole[r]; found {
		return skills
	}
	return []skill{}
}
