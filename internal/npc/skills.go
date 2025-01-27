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
	skillAdmin               = newSkill("Admin")
	skillAdvocate            = newSkill("Advocate")
	skillArtActing           = newSpecialization("Art", "Acting")
	skillArtInstrument       = newSpecialization("Art", "Instrument")
	skillAstrogation         = newSkill("Astrogation")
	skillAthleticsDexterity  = newSpecialization("Athletics", "Dexterity")
	skillAthleticsStrength   = newSpecialization("Athletics", "Strength")
	skillBiology             = newSpecialization("Science", "Biology")
	skillBroker              = newSkill("Broker")
	skillCarouse             = newSkill("Carouse")
	skillChemical            = newSpecialization("Science", "Chemical")
	skillCommunications      = newSpecialization("Electronics", "Communications")
	skillComputers           = newSpecialization("Electronics", "Computers")
	skillDeception           = newSkill("Deception")
	skillDiplomat            = newSkill("Diplomat")
	skillDriveAny            = newSpecialization("Drive", placeholderForSpecialization)
	skillDriveGrav           = newSpecialization("Drive", "Grav")
	skillElectronicSensors   = newSpecialization("Electronics", "Sensors")
	skillEngineerAny         = newSpecialization("Engineer", placeholderForSpecialization)
	skillEngineerJDrive      = newSpecialization("Engineer", "Jump Drive")
	skillEngineerLifeSupport = newSpecialization("Engineer", "Life Support")
	skillEngineerMDrive      = newSpecialization("Engineer", "Manoeuvre Drive")
	skillEngineerPower       = newSpecialization("Engineer", "Power")
	skillExplosives          = newSkill("Explosives")
	skillFlyer               = newSpecialization("Flyer", placeholderForSpecialization)
	skillGunnerAny           = newSpecialization("Gunner", placeholderForSpecialization)
	skillGunnerTurrets       = newSpecialization("Gunner", "Turrets")
	skillGunnerScreens       = newSpecialization("Gunner", "Screens")
	skillGunCombat           = newSpecialization("Gun Combat", placeholderForSpecialization)
	skillHeavyWeapons        = newSpecialization("Heavy Weapons", placeholderForSpecialization)
	skillInvestigate         = newSkill("Investigate")
	skillLanguage            = newSpecialization("Language", placeholderForSpecialization)
	skillLeadership          = newSkill("Leadership")
	skillMechanic            = newSkill("Mechanic")
	skillMedic               = newSkill("Medic")
	skillMeleeBlade          = newSpecialization("Melee", "Blade")
	skillMeleeUnarmed        = newSpecialization("Melee", "Unarmed")
	skillNavigation          = newSkill("Navigation")
	skillPersuade            = newSkill("Persuade")
	skillPilotAny            = newSpecialization("Pilot", placeholderForSpecialization)
	skillRecon               = newSkill("Recon")
	skillSmallCraft          = newSpecialization("Pilot", "Small craft")
	skillSpacecraft          = newSpecialization("Pilot", "Spacecraft")
	skillStealth             = newSkill("Stealth")
	skillSteward             = newSkill("Steward")
	skillStreetwise          = newSkill("Streetwise")
	skillSurvival            = newSkill("Survival")
	skillTacticsAny          = newSpecialization("Tactics", placeholderForSpecialization)
	skillVaccSuit            = newSkill("Vacc Suit")
)

var skillsByRole = map[Role][]skill{
	RolePilot: {
		skillSpacecraft,
		skillAstrogation,
		skillElectronicSensors,
		skillGunnerAny,
		skillMechanic,
		skillSmallCraft,
		skillLeadership,
		skillVaccSuit,
		skillCommunications,
		skillDriveGrav,
		skillSurvival,
		skillRecon,
		skillFlyer,
	},
	RoleNavigator: {
		skillAstrogation,
		skillElectronicSensors,
		skillSpacecraft,
		skillComputers,
		skillSurvival,
		skillNavigation,
		skillMechanic,
		skillLeadership,
		skillTacticsAny,
		skillEngineerAny,
		skillVaccSuit,
		skillRecon,
	},
	RoleEngineer: {
		skillEngineerMDrive,
		skillMechanic,
		skillEngineerPower,
		skillComputers,
		skillEngineerJDrive,
		skillEngineerLifeSupport,
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
		skillDiplomat,
	},
	RoleMedic: {
		skillMedic,
		skillBiology,
		skillChemical,
		skillDeception,
		skillInvestigate,
		skillDiplomat,
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
		skillStealth,
	},
	RoleGunner: {
		skillGunnerTurrets,
		skillElectronicSensors,
		skillGunnerScreens,
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
		skillGunnerAny,
		skillMedic,
		skillTacticsAny,
		skillGunCombat,
		skillNavigation,
		skillLeadership,
	},
	RoleTechnician: {
		skillMechanic,
		skillComputers,
		skillElectronicSensors,
		skillEngineerPower,
		skillEngineerMDrive,
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
		skillDiplomat,
		skillPersuade,
		skillAdvocate,
		skillElectronicSensors,
		skillComputers,
		skillDeception,
		skillSpacecraft,
		skillEngineerAny,
		skillRecon,
		skillMedic,
	},
	RoleDiplomat: {
		skillDiplomat,
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
		skillArtInstrument,
		skillPersuade,
		skillStealth,
		skillDeception,
		skillDiplomat,
		skillArtActing,
		skillComputers,
		skillElectronicSensors,
		skillLeadership,
		skillBroker,
		skillMeleeBlade,
		skillAdmin,
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
		skillDiplomat,
		skillDeception,
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
