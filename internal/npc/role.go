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

func (r Role) relevantSkills() []skill {
	if skills, found := skillsByRole[r]; found {
		return skills
	}
	return []skill{}
}

func (r Role) Skills(experience Experience) []string {
	if distribution, found := skillLevelDistributionByExperience[experience]; found {
		return distribution.DistributeLevels(r.relevantSkills())
	}
	return []string{}
}
