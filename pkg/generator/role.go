package generator

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
