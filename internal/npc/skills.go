package npc

import "fmt"

// For future reference, here is a list of skills: https://github.com/jacob7395/Traveller-Sheet/blob/master/SkillList.txt

const placeholderForSpecialization = "         "
const noSpecialization = ""

type skill struct {
	Name           string
	Specialization string
}

func newSkill(name string) skill {
	return skill{Name: name, Specialization: noSpecialization}
}

func newSpecialization(name, specialization string) skill {
	return skill{Name: name, Specialization: specialization}
}

func (s skill) StringLevel(level int) string {
	skillText := s.Name
	if s.Specialization != noSpecialization {
		skillText += fmt.Sprintf(" (%v)", s.Specialization)
	}
	return fmt.Sprintf("%v-%v", skillText, level)
}

var skillsByRole = map[Role][]skill{
	RolePilot: {
		newSpecialization("Pilot", "Spacecraft"),
		newSkill("Astrogation"),
		newSkill("Sensors"),
		newSkill("Gunnery"),
		newSkill("Mechanic"),
		newSkill("Communications"),
		newSkill("Leadership"),
		newSkill("Tactics"),
		newSkill("Small Craft"),
		newSkill("Engineering"),
		newSpecialization("Drive", "Grav"),
		newSkill("Survival"),
	},
	RoleNavigator: {
		newSkill("Astrogation"),
		newSkill("Sensors"),
		newSpecialization("Pilot", "Spacecraft"),
		newSpecialization("Electronics", "Computers"),
		newSkill("Survival"),
		newSpecialization("Electronic", "Sensors"),
		newSkill("Mechanic"),
		newSkill("Leadership"),
		newSkill("Tactics"),
		newSkill("Engineering"),
		newSpecialization("Navigation", "Surface"),
		newSkill("Liaison"),
	},
	RoleEngineer: {
		newSkill("Engineering"),
		newSkill("Mechanic"),
		newSpecialization("Electronics", "Power"),
		newSpecialization("Electronics", "Computers"),
		newSpecialization("Engineering", "Jump Drive"),
		newSpecialization("Engineering", "Life Support"),
		newSpecialization("Engineering", "Maneuver Drive"),
		newSkill("Sensors"),
		newSkill("Computer"),
		newSkill("Survival"),
		newSpecialization("Pilot", "Spacecraft"),
		newSkill("Leadership"),
		newSkill("Vacc Suit"),
	},
	RoleSteward: {
		newSkill("Steward"),
		newSkill("Carouse"),
		newSkill("Persuade"),
		newSkill("Broker"),
		newSkill("Admin"),
		newSpecialization("Electronic", "Computer"),
		newSpecialization("Language", placeholderForSpecialization),
		newSpecialization("Advocate", "Law"),
		newSkill("Leadership"),
		newSkill("Medic"),
		newSkill("Streetwise"),
		newSkill("Diplomacy"),
	},
	RoleMedic: {
		newSkill("Medic"),
		newSkill("Biology"),
		newSkill("Chemical"),
		newSpecialization("Electronics", "Medical"),
		newSkill("Diplomat"),
		newSkill("Persuade"),
		newSkill("Investigate"),
		newSkill("Broker"),
		newSpecialization("Electronics", "Computers"),
		newSkill("Admin"),
		newSkill("Sensors"),
		newSpecialization("Drive", placeholderForSpecialization),
	},
	RoleMarine: {
		newSkill("Gun Combat"),
		newSkill("Survival"),
		newSpecialization("Athletics", "Strength"),
		newSpecialization("Melee", "Unarmed"),
		newSkill("Heavy Weapons"),
		newSpecialization("Tactics", "Military"),
		newSkill("Recon"),
		newSkill("Sensors"),
		newSkill("Leadership"),
		newSkill("Medic"),
		newSpecialization("Pilot", "Ground"),
		newSkill("Communications"),
	},
	RoleGunner: {
		newSkill("Gunnery"),
		newSkill("Sensors"),
		newSkill("Tactics"),
		newSkill("Gun Combat"),
		newSkill("Leadership"),
		newSkill("Mechanic"),
		newSpecialization("Electronics", "Advanced Weapons"),
		newSkill("Computer"),
		newSkill("Heavy Weapons"),
		newSpecialization("Pilot", "Small Craft"),
		newSpecialization("Athletics", "Dexterity"),
		newSpecialization("Melee", "Blade"),
	},
	RoleScout: {
		newSkill("Survival"),
		newSkill("Recon"),
		newSpecialization("Pilot", "Small Craft"),
		newSkill("Astrogation"),
		newSkill("Sensors"),
		newSkill("Stealth"),
		newSkill("Gunnery"),
		newSkill("Medic"),
		newSkill("Tactics"),
		newSkill("Gun Combat"),
		newSkill("Navigation"),
		newSkill("Leadership"),
	},
	RoleTechnician: {
		newSkill("Mechanic"),
		newSpecialization("Electronics", "Computers"),
		newSpecialization("Electronics", "Sensors"),
		newSpecialization("Engineering", "Power"),
		newSpecialization("Engineering", "Maneuver Drive"),
		newSpecialization("Drive", placeholderForSpecialization),
		newSpecialization("Pilot", placeholderForSpecialization),
		newSkill("Vacc Suit"),
		newSkill("Recon"),
		newSkill("Athletics"),
		newSkill("Survival"),
		newSkill("Explosives"),
	},
	RoleLeader: {
		newSkill("Leadership"),
		newSkill("Tactics"),
		newSkill("Admin"),
		newSkill("Diplomat"),
		newSkill("Persuade"),
		newSkill("Advocate"),
		newSkill("Sensors"),
		newSpecialization("Electronics", "Computers"),
		newSpecialization("Pilot", "Spacecraft"),
		newSkill("Engineering"),
		newSkill("Medic"),
		newSkill("Recon"),
	},
	RoleDiplomat: {
		newSkill("Diplomat"),
		newSkill("Persuade"),
		newSkill("Advocate"),
		newSkill("Admin"),
		newSkill("Carouse"),
		newSkill("Streetwise"),
		newSkill("Linguistics"),
		newSkill("Leadership"),
		newSkill("Tactics"),
		newSpecialization("Electronics", "Computers"),
		newSkill("Sensors"),
		newSpecialization("Electronics", "Communications"),
	},
	RoleEntertainer: {
		newSkill("Carouse"),
		newSkill("Streetwise"),
		newSpecialization("Perform", placeholderForSpecialization),
		newSkill("Persuade"),
		newSkill("Stealth"),
		newSkill("Deception"),
		newSkill("Diplomat"),
		newSpecialization("Electronics", "Computers"),
		newSkill("Sensors"),
		newSkill("Leadership"),
		newSkill("Broker"),
		newSkill("Awareness"),
	},
	RoleTrader: {
		newSkill("Broker"),
		newSkill("Persuade"),
		newSkill("Admin"),
		newSkill("Advocate"),
		newSpecialization("Electronics", "Computers"),
		newSkill("Streetwise"),
		newSpecialization("Trade", placeholderForSpecialization),
		newSkill("Carouse"),
		newSkill("Diplomat"),
		newSkill("Awareness"),
		newSkill("Mechanic"),
		newSkill("Sensors"),
	},
	RoleThug: {
		newSpecialization("Melee", "Unarmed"),
		newSkill("Gun Combat"),
		newSpecialization("Melee", "Blade"),
		newSpecialization("Athletics", "Strength"),
		newSkill("Stealth"),
		newSkill("Streetwise"),
		newSkill("Carouse"),
		newSkill("Tactics"),
		newSkill("Awareness"),
		newSkill("Survival"),
		newSkill("Persuade"),
		newSkill("Grenades"),
		newSpecialization("Electronics", "Computers"),
	},
}
