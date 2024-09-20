package npc

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

func (r Role) RelevantSkills() []string {
	if skills, found := skillsByRole[r]; found {
		return skills
	}
	return []string{}
}

func (r Role) Skills(experience Experience) []string {
	if distribution, found := skillLevelDistributionByExperience[experience]; found {
		return distribution.DistributeLevels(r.RelevantSkills())
	}
	return []string{}
}

var skillsByRole = map[Role][]string{
	RolePilot: {
		"Pilot (Spacecraft)", "Astrogation", "Sensors", "Gunnery", "Mechanic", "Communications",
		"Leadership", "Tactics", "Small Craft", "Engineering", "Vehicle (Grav)", "Survival",
	},
	RoleNavigator: {
		"Astrogation", "Sensors", "Pilot (Spacecraft)", "Computers", "Survival", "Electronic (Sensors)",
		"Mechanic", "Leadership", "Tactics", "Engineering", "Navigation (Surface)", "Liaison",
	},
	RoleEngineer: {
		"Engineering", "Mechanic", "Electronics (Power)", "Electronics (Computers)", "Engineering (Jump Drive)",
		"Engineering (Life Support)", "Engineering (Maneuver Drive)", "Sensors", "Computer", "Survival",
		"Pilot (Spacecraft)", "Leadership",
	},
	RoleSteward: {
		"Steward", "Carouse", "Persuade", "Broker", "Admin", "Electronic (Computer)",
		"Language (Choose a specific language)", "Advocate (Law)", "Leadership", "Medic", "Streetwise", "Diplomacy",
	},
	RoleMedic: {
		"Medic", "Biology", "Chemical", "Electronics (Medical)", "Diplomat", "Persuade",
		"Investigate", "Broker", "Computers", "Admin", "Sensors", "Drive (Any)",
	},
	RoleMarine: {
		"Gunnery", "Survival", "Athletics (Strength)", "Melee (Unarmed)", "Tactics (Military)", "Recon",
		"Sensors", "Leadership", "Heavy Weapons", "Medic", "Piloting (Ground)", "Communications",
	},
	RoleGunner: {
		"Gunnery", "Sensors", "Tactics", "Gun Combat", "Leadership", "Mechanic",
		"Electronics (Advanced Weapons)", "Computer", "Heavy Weapons", "Pilot (Small Craft)",
		"Athletics (Dexterity)", "Melee (Blade)",
	},
	RoleScout: {
		"Survival", "Recon", "Pilot (Small Craft)", "Astrogation", "Sensors", "Stealth",
		"Gunnery", "Medic", "Tactics", "Gun Combat", "Navigation", "Leadership",
	},
	RoleTechnician: {
		"Mechanic", "Electronics (Computers)", "Engineering", "Sensors", "Computers",
		"Electronics (Power)", "Electronics (Robotics)", "Vehicle (Any)", "Piloting (Small Craft)",
		"Communications", "Athletics", "Programming",
	},
	RoleLeader: {
		"Leadership", "Tactics", "Admin", "Diplomat", "Persuade", "Advocate",
		"Sensors", "Computers", "Pilot (Spacecraft)", "Engineering", "Medic", "Recon",
	},
	RoleDiplomat: {
		"Diplomat", "Persuade", "Advocate", "Admin", "Carouse", "Streetwise",
		"Linguistics", "Leadership", "Tactics", "Computers", "Sensors", "Electronics (Communications)",
	},
	RoleEntertainer: {
		"Carouse", "Streetwise", "Perform (Any)", "Persuade", "Stealth", "Deception",
		"Diplomat", "Computers", "Sensors", "Leadership", "Broker", "Awareness",
	},
	RoleTrader: {
		"Broker", "Persuade", "Admin", "Advocate", "Computers", "Streetwise",
		"Trade (Any)", "Carouse", "Diplomat", "Awareness", "Mechanic", "Sensors",
	},
	RoleThug: {
		"Melee (Unarmed or Blade)", "Gun Combat", "Athletics (Strength)", "Stealth", "Streetwise", "Carouse",
		"Tactics", "Awareness", "Survival", "Persuade", "Grenades", "Computers",
	},
}
