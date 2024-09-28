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

var (
	skillAdmin                  = newSkill("Admin")
	skillAdvocate               = newSkill("Advocate")
	skillAstrogation            = newSkill("Astrogation")
	skillAthleticsDexterity     = newSpecialization("Athletics", "Dexterity")
	skillAthleticsStrength      = newSpecialization("Athletics", "Strength")
	skillBiology                = newSpecialization("Science", "Biology")
	skillBroker                 = newSkill("Broker")
	skillCarouse                = newSkill("Carouse")
	skillChemical               = newSpecialization("Science", "Chemical")
	skillCommunications         = newSpecialization("Electronics", "Communications")
	skillComputers              = newSpecialization("Electronics", "Computers")
	skillDeception              = newSkill("Deception")
	skillDiplomacy              = newSkill("Diplomacy")
	skillDriveAny               = newSpecialization("Drive", placeholderForSpecialization)
	skillDriveGrav              = newSpecialization("Drive", "Grav")
	skillElectronicPower        = newSpecialization("Electronic", "Power")
	skillElectronicSensors      = newSpecialization("Electronic", "Sensors")
	skillEngineeringAny         = newSpecialization("Engineering", placeholderForSpecialization)
	skillEngineeringJDrive      = newSpecialization("Engineering", "Jump Drive")
	skillEngineeringLifeSupport = newSpecialization("Engineering", "Life Support")
	skillEngineeringMDrive      = newSpecialization("Engineering", "Manoeuvre Drive")
	skillExplosives             = newSkill("Explosives")
	skillGunnery                = newSkill("Gunnery")
	skillGunCombat              = newSpecialization("Gun Combat", placeholderForSpecialization)
	skillHeavyWeapons           = newSkill("Heavy Weapons")
	skillInvestigate            = newSkill("Investigate")
	skillLanguage               = newSpecialization("Language", placeholderForSpecialization)
	skillLeadership             = newSkill("Leadership")
	skillMechanic               = newSkill("Mechanic")
	skillMedic                  = newSkill("Medic")
	skillMeleeBlade             = newSpecialization("Melee", "Blade")
	skillMeleeUnarmed           = newSpecialization("Melee", "Unarmed")
	skillNavigationSurface      = newSpecialization("Navigation", "Surface")
	skillPerform                = newSpecialization("Perform", placeholderForSpecialization)
	skillPersuade               = newSkill("Persuade")
	skillPilotAny               = newSpecialization("Pilot", placeholderForSpecialization)
	skillRecon                  = newSkill("Recon")
	skillSmallCraft             = newSpecialization("Pilot", "Small craft")
	skillSpacecraft             = newSpecialization("Pilot", "Spacecraft")
	skillStealth                = newSkill("Stealth")
	skillSteward                = newSkill("Steward")
	skillStreetwise             = newSkill("Streetwise")
	skillSurvival               = newSkill("Survival")
	skillTacticsAny             = newSpecialization("Tactics", placeholderForSpecialization)
	skillVaccSuit               = newSkill("Vacc Suit")
)

var skillsByRole = map[Role][]skill{
	RolePilot: {
		skillSpacecraft,
		skillAstrogation,
		skillElectronicSensors,
		skillGunnery,
		skillMechanic,
		skillSmallCraft,
		skillLeadership,
		skillVaccSuit,
		skillCommunications,
		skillDriveGrav,
		skillSurvival,
		skillRecon,
	},
	RoleNavigator: {
		skillAstrogation,
		skillElectronicSensors,
		skillSpacecraft,
		skillComputers,
		skillSurvival,
		skillNavigationSurface,
		skillMechanic,
		skillLeadership,
		skillTacticsAny,
		skillEngineeringAny,
		skillVaccSuit,
		skillRecon,
	},
	RoleEngineer: {
		skillEngineeringMDrive,
		skillMechanic,
		skillElectronicPower,
		skillComputers,
		skillEngineeringJDrive,
		skillEngineeringLifeSupport,
		skillElectronicSensors,
		skillSurvival,
		skillSmallCraft,
		skillLeadership,
		skillVaccSuit,
		skillRecon,
		skillDriveAny,
	},
	RoleSteward: {
		skillSteward,
		skillCarouse,
		skillPersuade,
		skillBroker,
		skillAdmin,
		skillComputers,
		skillLanguage,
		skillAdvocate,
		skillLeadership,
		skillMedic,
		skillStreetwise,
		skillDiplomacy,
	},
	RoleMedic: {
		skillMedic,
		skillBiology,
		skillChemical,
		skillDeception,
		skillInvestigate,
		skillDiplomacy,
		skillComputers,
		skillPersuade,
		skillAdmin,
		skillBroker,
		skillElectronicSensors,
		skillDriveAny,
		skillLeadership,
	},
	RoleMarine: {
		skillGunCombat,
		skillSurvival,
		skillAthleticsStrength,
		skillMeleeUnarmed,
		skillHeavyWeapons,
		skillTacticsAny,
		skillRecon,
		skillElectronicSensors,
		skillLeadership,
		skillMedic,
		skillDriveGrav,
		skillCommunications,
	},
	RoleGunner: {
		skillGunnery,
		skillElectronicSensors,
		skillTacticsAny,
		skillGunCombat,
		skillLeadership,
		skillMechanic,
		skillHeavyWeapons,
		skillExplosives,
		skillComputers,
		skillSmallCraft,
		skillAthleticsDexterity,
		skillMeleeBlade,
	},
	RoleScout: {
		skillSurvival,
		skillRecon,
		skillSmallCraft,
		skillAstrogation,
		skillElectronicSensors,
		skillStealth,
		skillGunnery,
		skillMedic,
		skillTacticsAny,
		skillGunCombat,
		skillNavigationSurface,
		skillLeadership,
	},
	RoleTechnician: {
		skillMechanic,
		skillComputers,
		skillElectronicSensors,
		skillElectronicPower,
		skillEngineeringMDrive,
		skillDriveAny,
		skillPilotAny,
		skillVaccSuit,
		skillRecon,
		skillAthleticsDexterity,
		skillSurvival,
		skillExplosives,
	},
	RoleLeader: {
		skillLeadership,
		skillTacticsAny,
		skillAdmin,
		skillDiplomacy,
		skillPersuade,
		skillAdvocate,
		skillElectronicSensors,
		skillComputers,
		skillSpacecraft,
		skillEngineeringAny,
		skillMedic,
		skillRecon,
	},
	RoleDiplomat: {
		skillDiplomacy,
		skillPersuade,
		skillAdvocate,
		skillAdmin,
		skillCarouse,
		skillSteward,
		skillStreetwise,
		skillLanguage,
		skillBroker,
		skillLeadership,
		skillCommunications,
		skillTacticsAny,
	},
	RoleEntertainer: {
		skillCarouse,
		skillStreetwise,
		skillPerform,
		skillPersuade,
		skillStealth,
		skillDeception,
		skillDiplomacy,
		skillComputers,
		skillElectronicSensors,
		skillLeadership,
		skillBroker,
		skillMeleeBlade,
	},
	RoleTrader: {
		skillBroker,
		skillPersuade,
		skillAdmin,
		skillAdvocate,
		skillComputers,
		skillStreetwise,
		skillGunCombat,
		skillComputers,
		skillDiplomacy,
		skillCarouse,
		skillCommunications,
		skillMechanic,
		skillElectronicSensors,
		skillLeadership,
	},
	RoleThug: {
		skillMeleeUnarmed,
		skillGunCombat,
		skillMeleeBlade,
		skillAthleticsStrength,
		skillStealth,
		skillStreetwise,
		skillCarouse,
		skillTacticsAny,
		skillStealth,
		skillSurvival,
		skillPersuade,
		skillExplosives,
		skillComputers,
	},
}
