package generator

import "github.com/carloscasalar/traveller-npc-generator/internal/npc"

//go:generate enumer -type=Role -output=role_auto.go -trimprefix=Role -transform=snake
type Role int

func (i Role) toNpcRole() npc.Role {
	switch i {
	case RolePilot:
		return npc.RolePilot
	case RoleNavigator:
		return npc.RoleNavigator
	case RoleEngineer:
		return npc.RoleEngineer
	case RoleSteward:
		return npc.RoleSteward
	case RoleMedic:
		return npc.RoleMedic
	case RoleMarine:
		return npc.RoleMarine
	case RoleGunner:
		return npc.RoleGunner
	case RoleScout:
		return npc.RoleScout
	case RoleTechnician:
		return npc.RoleTechnician
	case RoleLeader:
		return npc.RoleLeader
	case RoleDiplomat:
		return npc.RoleDiplomat
	case RoleEntertainer:
		return npc.RoleEntertainer
	case RoleTrader:
		return npc.RoleTrader
	case RoleThug:
		return npc.RoleThug
	default:
		return npc.Role(i)
	}
}

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
