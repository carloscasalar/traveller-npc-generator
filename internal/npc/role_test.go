package npc_test

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/internal/npc"
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestRole_Given_Recruit_NPC(t *testing.T) {
	roles := []npc.Role{
		npc.RolePilot,
		npc.RoleNavigator,
		npc.RoleEngineer,
		npc.RoleSteward,
		npc.RoleMedic,
		npc.RoleMarine,
		npc.RoleGunner,
		npc.RoleScout,
		npc.RoleTechnician,
		npc.RoleLeader,
		npc.RoleDiplomat,
		npc.RoleEntertainer,
		npc.RoleTrader,
		npc.RoleThug,
	}

	for _, role := range roles {
		t.Run(fmt.Sprintf("for the role %v should have 2 skills level 2, 2 skills level 1, 5 skills level 0", role.String()), func(t *testing.T) {
			skills := role.Skills(npc.ExperienceRecruit)

			assert.Len(t, getDuplicatedSkills(skills), 0)
			assert.Equal(t, 0, numberOfSkillsAtLevel(skills, 3))
			assert.Equal(t, 0, numberOfSkillsAtLevel(skills, 2))
			assert.Equal(t, 0, numberOfSkillsAtLevel(skills, 1))
			assert.Equal(t, 4, numberOfSkillsAtLevel(skills, 0))
		})
	}
}

func TestRole_Given_Rookie_NPC(t *testing.T) {
	roles := []npc.Role{
		npc.RolePilot,
		npc.RoleNavigator,
		npc.RoleEngineer,
		npc.RoleSteward,
		npc.RoleMedic,
		npc.RoleMarine,
		npc.RoleGunner,
		npc.RoleScout,
		npc.RoleTechnician,
		npc.RoleLeader,
		npc.RoleDiplomat,
		npc.RoleEntertainer,
		npc.RoleTrader,
		npc.RoleThug,
	}

	for _, role := range roles {
		t.Run(fmt.Sprintf("for the role %v should have 2 skills level 2, 2 skills level 1, 5 skills level 0", role.String()), func(t *testing.T) {
			skills := role.Skills(npc.ExperienceRookie)

			assert.Len(t, getDuplicatedSkills(skills), 0)
			assert.Equal(t, 0, numberOfSkillsAtLevel(skills, 3))
			assert.Equal(t, 0, numberOfSkillsAtLevel(skills, 2))
			assert.Equal(t, 2, numberOfSkillsAtLevel(skills, 1))
			assert.Equal(t, 4, numberOfSkillsAtLevel(skills, 0))
		})
	}
}

func TestRole_Given_Intermediate_NPC(t *testing.T) {
	roles := []npc.Role{
		npc.RolePilot,
		npc.RoleNavigator,
		npc.RoleEngineer,
		npc.RoleSteward,
		npc.RoleMedic,
		npc.RoleMarine,
		npc.RoleGunner,
		npc.RoleScout,
		npc.RoleTechnician,
		npc.RoleLeader,
		npc.RoleDiplomat,
		npc.RoleEntertainer,
		npc.RoleTrader,
		npc.RoleThug,
	}

	for _, role := range roles {
		t.Run(fmt.Sprintf("for the role %v should have 1 skills level 2, 2 skills level 1, 4 skills level 0", role.String()), func(t *testing.T) {
			skills := role.Skills(npc.ExperienceIntermediate)

			assert.Len(t, getDuplicatedSkills(skills), 0)
			assert.Equal(t, 0, numberOfSkillsAtLevel(skills, 3))
			assert.Equal(t, 1, numberOfSkillsAtLevel(skills, 2))
			assert.Equal(t, 2, numberOfSkillsAtLevel(skills, 1))
			assert.Equal(t, 4, numberOfSkillsAtLevel(skills, 0))
		})
	}
}

func TestRole_Given_Regular_NPC(t *testing.T) {
	roles := []npc.Role{
		npc.RolePilot,
		npc.RoleNavigator,
		npc.RoleEngineer,
		npc.RoleSteward,
		npc.RoleMedic,
		npc.RoleMarine,
		npc.RoleGunner,
		npc.RoleScout,
		npc.RoleTechnician,
		npc.RoleLeader,
		npc.RoleDiplomat,
		npc.RoleEntertainer,
		npc.RoleTrader,
		npc.RoleThug,
	}

	for _, role := range roles {
		t.Run(fmt.Sprintf("for the role %v should have 2 skills level 2, 2 skills level 1, 5 skills level 0", role.String()), func(t *testing.T) {
			skills := role.Skills(npc.ExperienceRegular)

			assert.Len(t, getDuplicatedSkills(skills), 0)
			assert.Equal(t, 0, numberOfSkillsAtLevel(skills, 3))
			assert.Equal(t, 2, numberOfSkillsAtLevel(skills, 2))
			assert.Equal(t, 2, numberOfSkillsAtLevel(skills, 1))
			assert.Equal(t, 5, numberOfSkillsAtLevel(skills, 0))
		})
	}
}

func TestRole_Given_Veteran_NPC(t *testing.T) {
	roles := []npc.Role{
		npc.RolePilot,
		npc.RoleNavigator,
		npc.RoleEngineer,
		npc.RoleSteward,
		npc.RoleMedic,
		npc.RoleMarine,
		npc.RoleGunner,
		npc.RoleScout,
		npc.RoleTechnician,
		npc.RoleLeader,
		npc.RoleDiplomat,
		npc.RoleEntertainer,
		npc.RoleTrader,
		npc.RoleThug,
	}

	for _, role := range roles {
		t.Run(fmt.Sprintf("for the role %v should have 3 skills level 2, 2 skills level 1, 5 skills level 0", role.String()), func(t *testing.T) {
			skills := role.Skills(npc.ExperienceVeteran)

			assert.Len(t, getDuplicatedSkills(skills), 0)
			assert.Equal(t, 0, numberOfSkillsAtLevel(skills, 3))
			assert.Equal(t, 3, numberOfSkillsAtLevel(skills, 2))
			assert.Equal(t, 2, numberOfSkillsAtLevel(skills, 1))
			assert.Equal(t, 5, numberOfSkillsAtLevel(skills, 0))
		})
	}
}

func TestRole_Given_Elite_NPC(t *testing.T) {
	roles := []npc.Role{
		npc.RolePilot,
		npc.RoleNavigator,
		npc.RoleEngineer,
		npc.RoleSteward,
		npc.RoleMedic,
		npc.RoleMarine,
		npc.RoleGunner,
		npc.RoleScout,
		npc.RoleTechnician,
		npc.RoleLeader,
		npc.RoleDiplomat,
		npc.RoleEntertainer,
		npc.RoleTrader,
		npc.RoleThug,
	}

	for _, role := range roles {
		t.Run(fmt.Sprintf("for the role %v should have 3 skills level 2, 2 skills level 1, 5 skills level 0", role.String()), func(t *testing.T) {
			skills := role.Skills(npc.ExperienceElite)
			assert.Len(t, getDuplicatedSkills(skills), 0)
			assert.Equal(t, 1, numberOfSkillsAtLevel(skills, 3))
			assert.Equal(t, 2, numberOfSkillsAtLevel(skills, 2))
			assert.Equal(t, 3, numberOfSkillsAtLevel(skills, 1))
			assert.Equal(t, 6, numberOfSkillsAtLevel(skills, 0))
		})
	}
}

func TestRole_RandomCharacteristic_given_an_exceptional_pilot(t *testing.T) {
	role := npc.RolePilot
	category := npc.CategoryExceptional

	characteristics := role.RandomCharacteristic(category)

	// Two highest characteristics should be INT and DEX
	intAndDex := []int{characteristics[npc.INT], characteristics[npc.DEX]}
	slices.Sort(intAndDex)
	slices.Reverse(intAndDex)
	assert.Equal(t, []int{11, 10}, intAndDex)

	// Two medium characteristics should be EDU and STR
	eduAndStr := []int{characteristics[npc.EDU], characteristics[npc.STR]}
	slices.Sort(eduAndStr)
	slices.Reverse(eduAndStr)
	assert.Equal(t, []int{9, 9}, eduAndStr)

	// Two lowest characteristics should be END and SOC
	endAndSoc := []int{characteristics[npc.END], characteristics[npc.SOC]}
	slices.Sort(endAndSoc)
	slices.Reverse(endAndSoc)
	assert.Equal(t, []int{8, 7}, endAndSoc)
}

func getLevel(skill string) int {
	return int(skill[len(skill)-1] - '0')
}

func getSkill(skill string) string {
	return skill[:len(skill)-2]
}

func numberOfSkillsAtLevel(skills []string, level int) int {
	count := 0
	for _, skill := range skills {
		if getLevel(skill) == level {
			count++
		}
	}
	return count
}

func getDuplicatedSkills(skillsAndLevels []string) []string {
	uniqueSkills := make(map[string]bool)
	duplicatedSkills := make([]string, 0)
	for _, s := range skillsAndLevels {
		skill := getSkill(s)
		if _, found := uniqueSkills[skill]; found {
			duplicatedSkills = append(duplicatedSkills, skill)
		}
		uniqueSkills[skill] = true
	}
	return duplicatedSkills
}
