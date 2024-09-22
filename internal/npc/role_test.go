package npc_test

import (
	"fmt"
	"github.com/carloscasalar/traveller-npc-generator/internal/npc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRole_Given_Recruit_NPC(t *testing.T) {
	tests := []struct {
		role                npc.Role
		expectedSkillLevels []string
	}{
		{npc.RolePilot, []string{"Pilot-0", "Astrogation-0", "Sensors-0", "Gunnery-0"}},
		{npc.RoleNavigator, []string{"Astrogation-0", "Sensors-0", "Pilot-0", "Computers-0"}},
		{npc.RoleEngineer, []string{"Engineering-0", "Mechanic-0", "Electronics-0", "Sensors-0"}},
		{npc.RoleSteward, []string{"Steward-0", "Carouse-0", "Persuade-0", "Broker-0"}},
		{npc.RoleMedic, []string{"Medic-0", "Biology-0", "Chemical-0", "Electronics-0"}},
		{npc.RoleMarine, []string{"Gunnery-0", "Survival-0", "Athletics-0", "Melee-0"}},
		{npc.RoleGunner, []string{"Gunnery-0", "Sensors-0", "Tactics-0", "Gun Combat-0"}},
		{npc.RoleScout, []string{"Survival-0", "Recon-0", "Pilot-0", "Astrogation-0"}},
		{npc.RoleTechnician, []string{"Mechanic-0", "Electronics-0", "Engineering-0", "Sensors-0"}},
		{npc.RoleLeader, []string{"Leadership-0", "Tactics-0", "Admin-0", "Diplomat-0"}},
		{npc.RoleDiplomat, []string{"Diplomat-0", "Persuade-0", "Advocate-0", "Admin-0"}},
		{npc.RoleEntertainer, []string{"Carouse-0", "Streetwise-0", "Perform-0", "Persuade-0"}},
		{npc.RoleTrader, []string{"Broker-0", "Persuade-0", "Admin-0", "Advocate-0"}},
		{npc.RoleThug, []string{"Melee-0", "Gun Combat-0", "Athletics-0", "Stealth-0"}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for the role %v should have 4 skills at level 0", tt.role.String()), func(t *testing.T) {
			skills := tt.role.Skills(npc.ExperienceRecruit)
			assert.Equal(t, tt.expectedSkillLevels, skills)
		})
	}
}

func TestRole_Given_Rookie_NPC(t *testing.T) {
	tests := []struct {
		role                npc.Role
		expectedSkillLevels []string
	}{
		{npc.RolePilot, []string{"Pilot (Spacecraft)-1", "Astrogation-1", "Sensors-0", "Gunnery-0", "Mechanic-0", "Communications-0"}},
		{npc.RoleNavigator, []string{"Astrogation-1", "Sensors-1", "Pilot-0", "Computers-0", "Survival-0", "Electronic-0"}},
		{npc.RoleEngineer, []string{"Engineering-1", "Mechanic-1", "Electronics-0", "Sensors-0", "Computer-0", "Survival-0"}},
		{npc.RoleSteward, []string{"Steward-1", "Carouse-1", "Persuade-0", "Broker-0", "Admin-0", "Electronic-0"}},
		{npc.RoleMedic, []string{"Medic-1", "Biology-1", "Chemical-0", "Electronics-0", "Diplomat-0", "Persuade-0"}},
		{npc.RoleMarine, []string{"Gunnery-1", "Survival-1", "Athletics-0", "Melee-0", "Tactics-0", "Recon-0"}},
		{npc.RoleGunner, []string{"Gunnery-1", "Sensors-1", "Tactics-0", "Gun Combat-0", "Leadership-0", "Mechanic-0"}},
		{npc.RoleScout, []string{"Survival-1", "Recon-1", "Pilot-0", "Astrogation-0", "Sensors-0", "Stealth-0"}},
		{npc.RoleTechnician, []string{"Mechanic-1", "Electronics (Computers)-1", "Engineering-0", "Sensors-0", "Computers-0", "Vehicle-0"}},
		{npc.RoleLeader, []string{"Leadership-1", "Tactics-1", "Admin-0", "Diplomat-0", "Persuade-0", "Advocate-0"}},
		{npc.RoleDiplomat, []string{"Diplomat-1", "Persuade-1", "Advocate-0", "Admin-0", "Carouse-0", "Streetwise-0"}},
		{npc.RoleEntertainer, []string{"Carouse-1", "Streetwise-1", "Perform-0", "Persuade-0", "Stealth-0", "Deception-0"}},
		{npc.RoleTrader, []string{"Broker-1", "Persuade-1", "Admin-0", "Advocate-0", "Computers-0", "Streetwise-0"}},
		{npc.RoleThug, []string{"Melee (Unarmed)-1", "Gun Combat-1", "Athletics-0", "Stealth-0", "Streetwise-0", "Carouse-0"}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for the role %v should have 2 sikills level 1, 4 skills level 0", tt.role.String()), func(t *testing.T) {
			skills := tt.role.Skills(npc.ExperienceRookie)
			assert.Equal(t, tt.expectedSkillLevels, skills)
		})
	}
}

func TestRole_Given_Intermediate_NPC(t *testing.T) {
	tests := []struct {
		role                npc.Role
		expectedSkillLevels []string
	}{
		{npc.RolePilot, []string{"Pilot (Spacecraft)-2", "Astrogation-1", "Sensors-1", "Gunnery-0", "Mechanic-0", "Communications-0", "Leadership-0"}},
		{npc.RoleNavigator, []string{"Astrogation-2", "Sensors-1", "Pilot (Spacecraft)-1", "Computers-0", "Survival-0", "Electronic-0", "Mechanic-0"}},
		{npc.RoleEngineer, []string{"Engineering-2", "Mechanic-1", "Electronics (Power)-1", "Sensors-0", "Computer-0", "Survival-0", "Pilot-0"}},
		{npc.RoleSteward, []string{"Steward-2", "Carouse-1", "Persuade-1", "Broker-0", "Admin-0", "Electronic-0", "Language-0"}},
		{npc.RoleMedic, []string{"Medic-2", "Biology-1", "Chemical-1", "Electronics-0", "Diplomat-0", "Persuade-0", "Investigate-0"}},
		{npc.RoleMarine, []string{"Gunnery-2", "Survival-1", "Athletics (Strength)-1", "Melee-0", "Tactics-0", "Recon-0", "Sensors-0"}},
		{npc.RoleGunner, []string{"Gunnery-2", "Sensors-1", "Tactics-1", "Gun Combat-0", "Leadership-0", "Mechanic-0", "Electronics-0"}},
		{npc.RoleScout, []string{"Survival-2", "Recon-1", "Pilot (Small Craft)-1", "Astrogation-0", "Sensors-0", "Stealth-0", "Gunnery-0"}},
		{npc.RoleTechnician, []string{"Mechanic-2", "Electronics (Computers)-1", "Engineering-1", "Sensors-0", "Computers-0", "Vehicle-0", "Piloting-0"}},
		{npc.RoleLeader, []string{"Leadership-2", "Tactics-1", "Admin-1", "Diplomat-0", "Persuade-0", "Advocate-0", "Sensors-0"}},
		{npc.RoleDiplomat, []string{"Diplomat-2", "Persuade-1", "Advocate-1", "Admin-0", "Carouse-0", "Streetwise-0", "Linguistics-0"}},
		{npc.RoleEntertainer, []string{"Carouse-2", "Streetwise-1", "Perform (         )-1", "Persuade-0", "Stealth-0", "Deception-0", "Diplomat-0"}},
		{npc.RoleTrader, []string{"Broker-2", "Persuade-1", "Admin-1", "Advocate-0", "Computers-0", "Streetwise-0", "Trade-0"}},
		{npc.RoleThug, []string{"Melee (Unarmed)-2", "Gun Combat-1", "Melee (Blade)-1", "Athletics-0", "Stealth-0", "Streetwise-0", "Carouse-0"}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for the role %v should have 1 skill level 2, 2 skills level 1, 4 skills level 0", tt.role.String()), func(t *testing.T) {
			skills := tt.role.Skills(npc.ExperienceIntermediate)
			assert.Equal(t, tt.expectedSkillLevels, skills)
		})
	}
}

func TestRole_Given_Regular_NPC(t *testing.T) {
	tests := []struct {
		role                npc.Role
		expectedSkillLevels []string
	}{
		{npc.RolePilot, []string{"Pilot (Spacecraft)-2", "Astrogation-2", "Sensors-1", "Gunnery-1", "Mechanic-0", "Communications-0", "Leadership-0", "Tactics-0", "Small Craft-0"}},
		{npc.RoleNavigator, []string{"Astrogation-2", "Sensors-2", "Pilot (Spacecraft)-1", "Computers-1", "Survival-0", "Electronic-0", "Mechanic-0", "Leadership-0", "Tactics-0"}},
		{npc.RoleEngineer, []string{"Engineering-2", "Mechanic-2", "Electronics (Power)-1", "Electronics (Computers)-1", "Sensors-0", "Computer-0", "Survival-0", "Pilot-0", "Leadership-0"}},
		{npc.RoleSteward, []string{"Steward-2", "Carouse-2", "Persuade-1", "Broker-1", "Admin-0", "Electronic-0", "Language-0", "Advocate-0", "Leadership-0"}},
		{npc.RoleMedic, []string{"Medic-2", "Biology-2", "Chemical-1", "Electronics (Medical)-1", "Diplomat-0", "Persuade-0", "Investigate-0", "Broker-0", "Computers-0"}},
		{npc.RoleMarine, []string{"Gunnery-2", "Survival-2", "Athletics (Strength)-1", "Melee (Unarmed)-1", "Tactics-0", "Recon-0", "Sensors-0", "Leadership-0", "Heavy Weapons-0"}},
		{npc.RoleGunner, []string{"Gunnery-2", "Sensors-2", "Tactics-1", "Gun Combat-1", "Leadership-0", "Mechanic-0", "Electronics-0", "Computer-0", "Heavy Weapons-0"}},
		{npc.RoleScout, []string{"Survival-2", "Recon-2", "Pilot (Small Craft)-1", "Astrogation-1", "Sensors-0", "Stealth-0", "Gunnery-0", "Medic-0", "Tactics-0"}},
		{npc.RoleTechnician, []string{"Mechanic-2", "Electronics (Computers)-2", "Engineering-1", "Sensors-1", "Computers-0", "Vehicle-0", "Piloting-0", "Communications-0", "Athletics-0"}},
		{npc.RoleLeader, []string{"Leadership-2", "Tactics-2", "Admin-1", "Diplomat-1", "Persuade-0", "Advocate-0", "Sensors-0", "Computers-0", "Pilot-0"}},
		{npc.RoleDiplomat, []string{"Diplomat-2", "Persuade-2", "Advocate-1", "Admin-1", "Carouse-0", "Streetwise-0", "Linguistics-0", "Leadership-0", "Tactics-0"}},
		{npc.RoleEntertainer, []string{"Carouse-2", "Streetwise-2", "Perform (         )-1", "Persuade-1", "Stealth-0", "Deception-0", "Diplomat-0", "Computers-0", "Sensors-0"}},
		{npc.RoleTrader, []string{"Broker-2", "Persuade-2", "Admin-1", "Advocate-1", "Computers-0", "Streetwise-0", "Trade-0", "Carouse-0", "Diplomat-0"}},
		{npc.RoleThug, []string{"Melee (Unarmed)-2", "Gun Combat-2", "Melee (Blade)-1", "Athletics (Strength)-1", "Stealth-0", "Streetwise-0", "Carouse-0", "Tactics-0", "Awareness-0"}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for the role %v should have 2 skills level 2, 2 skills level 1, 5 skills level 0", tt.role.String()), func(t *testing.T) {
			skills := tt.role.Skills(npc.ExperienceRegular)
			assert.Equal(t, tt.expectedSkillLevels, skills)
		})
	}
}

func TestRole_Given_Veteran_NPC(t *testing.T) {
	tests := []struct {
		role                npc.Role
		expectedSkillLevels []string
	}{
		{npc.RolePilot, []string{"Pilot (Spacecraft)-2", "Astrogation-2", "Sensors-2", "Gunnery-1", "Mechanic-1", "Communications-0", "Leadership-0", "Tactics-0", "Small Craft-0", "Engineering-0"}},
		{npc.RoleNavigator, []string{"Astrogation-2", "Sensors-2", "Pilot (Spacecraft)-2", "Computers-1", "Survival-1", "Electronic-0", "Mechanic-0", "Leadership-0", "Tactics-0", "Engineering-0"}},
		{npc.RoleEngineer, []string{"Engineering-2", "Mechanic-2", "Electronics (Power)-2", "Electronics (Computers)-1", "Engineering (Jump Drive)-1", "Sensors-0", "Computer-0", "Survival-0", "Pilot-0", "Leadership-0"}},
		{npc.RoleSteward, []string{"Steward-2", "Carouse-2", "Persuade-2", "Broker-1", "Admin-1", "Electronic-0", "Language-0", "Advocate-0", "Leadership-0", "Medic-0"}},
		{npc.RoleMedic, []string{"Medic-2", "Biology-2", "Chemical-2", "Electronics (Medical)-1", "Diplomat-1", "Persuade-0", "Investigate-0", "Broker-0", "Computers-0", "Admin-0"}},
		{npc.RoleMarine, []string{"Gunnery-2", "Survival-2", "Athletics (Strength)-2", "Melee (Unarmed)-1", "Tactics (Military)-1", "Recon-0", "Sensors-0", "Leadership-0", "Heavy Weapons-0", "Medic-0"}},
		{npc.RoleGunner, []string{"Gunnery-2", "Sensors-2", "Tactics-2", "Gun Combat-1", "Leadership-1", "Mechanic-0", "Electronics-0", "Computer-0", "Heavy Weapons-0", "Pilot-0"}},
		{npc.RoleScout, []string{"Survival-2", "Recon-2", "Pilot (Small Craft)-2", "Astrogation-1", "Sensors-1", "Stealth-0", "Gunnery-0", "Medic-0", "Tactics-0", "Gun Combat-0"}},
		{npc.RoleTechnician, []string{"Mechanic-2", "Electronics (Computers)-2", "Engineering-2", "Sensors-1", "Computers-1", "Vehicle-0", "Piloting-0", "Communications-0", "Athletics-0", "Programming-0"}},
		{npc.RoleLeader, []string{"Leadership-2", "Tactics-2", "Admin-2", "Diplomat-1", "Persuade-1", "Advocate-0", "Sensors-0", "Computers-0", "Pilot-0", "Engineering-0"}},
		{npc.RoleDiplomat, []string{"Diplomat-2", "Persuade-2", "Advocate-2", "Admin-1", "Carouse-1", "Streetwise-0", "Linguistics-0", "Leadership-0", "Tactics-0", "Computers-0"}},
		{npc.RoleEntertainer, []string{"Carouse-2", "Streetwise-2", "Perform (         )-2", "Persuade-1", "Stealth-1", "Deception-0", "Diplomat-0", "Computers-0", "Sensors-0", "Leadership-0"}},
		{npc.RoleTrader, []string{"Broker-2", "Persuade-2", "Admin-2", "Advocate-1", "Computers-1", "Streetwise-0", "Trade-0", "Carouse-0", "Diplomat-0", "Awareness-0"}},
		{npc.RoleThug, []string{"Melee (Unarmed)-2", "Gun Combat-2", "Melee (Blade)-2", "Athletics (Strength)-1", "Stealth-1", "Streetwise-0", "Carouse-0", "Tactics-0", "Awareness-0", "Survival-0"}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for the role %v should have 3 skills level 2, 2 skills level 1, 5 skills level 0", tt.role.String()), func(t *testing.T) {
			skills := tt.role.Skills(npc.ExperienceVeteran)
			assert.Equal(t, tt.expectedSkillLevels, skills)
		})
	}
}

func TestRole_Given_Elite_NPC(t *testing.T) {
	tests := []struct {
		role                npc.Role
		expectedSkillLevels []string
	}{
		{npc.RolePilot, []string{"Pilot (Spacecraft)-3", "Astrogation-2", "Sensors-2", "Gunnery-1", "Mechanic-1", "Communications-1", "Leadership-0", "Tactics-0", "Small Craft-0", "Engineering-0", "Vehicle-0", "Survival-0"}},
		{npc.RoleNavigator, []string{"Astrogation-3", "Sensors-2", "Pilot (Spacecraft)-2", "Computers-1", "Survival-1", "Electronic (Sensors)-1", "Mechanic-0", "Leadership-0", "Tactics-0", "Engineering-0", "Navigation-0", "Liaison-0"}},
		{npc.RoleEngineer, []string{"Engineering-3", "Mechanic-2", "Electronics (Power)-2", "Electronics (Computers)-1", "Engineering (Jump Drive)-1", "Engineering (Life Support)-1", "Sensors-0", "Computer-0", "Survival-0", "Pilot-0", "Leadership-0"}},
		{npc.RoleSteward, []string{"Steward-3", "Carouse-2", "Persuade-2", "Broker-1", "Admin-1", "Electronic (Computer)-1", "Language-0", "Advocate-0", "Leadership-0", "Medic-0", "Streetwise-0", "Diplomacy-0"}},
		{npc.RoleMedic, []string{"Medic-3", "Biology-2", "Chemical-2", "Electronics (Medical)-1", "Diplomat-1", "Persuade-1", "Investigate-0", "Broker-0", "Computers-0", "Admin-0", "Sensors-0", "Drive-0"}},
		{npc.RoleMarine, []string{"Gunnery-3", "Survival-2", "Athletics (Strength)-2", "Melee (Unarmed)-1", "Tactics (Military)-1", "Recon-1", "Sensors-0", "Leadership-0", "Heavy Weapons-0", "Medic-0", "Piloting-0", "Communications-0"}},
		{npc.RoleGunner, []string{"Gunnery-3", "Sensors-2", "Tactics-2", "Gun Combat-1", "Leadership-1", "Mechanic-1", "Electronics-0", "Computer-0", "Heavy Weapons-0", "Pilot-0", "Athletics-0", "Melee-0"}},
		{npc.RoleScout, []string{"Survival-3", "Recon-2", "Pilot (Small Craft)-2", "Astrogation-1", "Sensors-1", "Stealth-1", "Gunnery-0", "Medic-0", "Tactics-0", "Gun Combat-0", "Navigation-0", "Leadership-0"}},
		{npc.RoleTechnician, []string{"Mechanic-3", "Electronics (Computers)-2", "Engineering-2", "Sensors-1", "Computers-1", "Electronics (Power)-1", "Vehicle-0", "Piloting-0", "Communications-0", "Athletics-0", "Programming-0"}},
		{npc.RoleLeader, []string{"Leadership-3", "Tactics-2", "Admin-2", "Diplomat-1", "Persuade-1", "Advocate-1", "Sensors-0", "Computers-0", "Pilot-0", "Engineering-0", "Medic-0", "Recon-0"}},
		{npc.RoleDiplomat, []string{"Diplomat-3", "Persuade-2", "Advocate-2", "Admin-1", "Carouse-1", "Streetwise-1", "Linguistics-0", "Leadership-0", "Tactics-0", "Computers-0", "Sensors-0", "Electronics-0"}},
		{npc.RoleEntertainer, []string{"Carouse-3", "Streetwise-2", "Perform (         )-2", "Persuade-1", "Stealth-1", "Deception-1", "Diplomat-0", "Computers-0", "Sensors-0", "Leadership-0", "Broker-0", "Awareness-0"}},
		{npc.RoleTrader, []string{"Broker-3", "Persuade-2", "Admin-2", "Advocate-1", "Computers-1", "Streetwise-1", "Trade-0", "Carouse-0", "Diplomat-0", "Awareness-0", "Mechanic-0", "Sensors-0"}},
		{npc.RoleThug, []string{"Melee (Unarmed)-3", "Gun Combat-2", "Melee (Blade)-2", "Athletics (Strength)-1", "Stealth-1", "Streetwise-1", "Carouse-0", "Tactics-0", "Awareness-0", "Survival-0", "Persuade-0", "Grenades-0"}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("for the role %v should have 3 skills level 2, 2 skills level 1, 5 skills level 0", tt.role.String()), func(t *testing.T) {
			skills := tt.role.Skills(npc.ExperienceElite)
			assert.Equal(t, tt.expectedSkillLevels, skills)
		})
	}
}
