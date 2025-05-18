package generator

import (

	// math/rand, os, time, gopkg.in/yaml.v3 are no longer directly needed here

	"github.com/carloscasalar/traveller-npc-generator/internal/gear"
)

// GearItem represents a generic piece of equipment for the public API.
type GearItem struct {
	Name             string   `json:"name"` // JSON tags for consistency if this struct is ever serialized by clients
	Type             string   `json:"type"`
	WeaponType       string   `json:"weapon_type,omitempty"`
	ArmorValue       int      `json:"armor_value,omitempty"`
	TechLevel        int      `json:"tech_level"`
	Damage           string   `json:"damage,omitempty"`
	RangeMeters      int      `json:"range_meters,omitempty"`
	WeightKg         float64  `json:"weight_kg"`
	MagazineCapacity int      `json:"magazine_capacity,omitempty"`
	AmmoType         string   `json:"ammo_type,omitempty"`
	Legality         string   `json:"legality"`
	CostCredits      int      `json:"cost_credits"`
	Tags             []string `json:"tags"`
}

// EquipmentSet holds all gear for an NPC for the public API.
type EquipmentSet struct {
	Armor   []GearItem
	Weapons []GearItem
	Tools   []GearItem
	Misc    []GearItem
}

// InitGearData initializes the gear data by calling the internal gear package.
func InitGearData() error {
	return gear.Init()
}

// --- Mappers ---

func mapInternalItemToGeneratorItem(internalItem gear.Item) GearItem {
	return GearItem{
		Name:             internalItem.Name,
		Type:             internalItem.Type,
		WeaponType:       internalItem.WeaponType,
		ArmorValue:       internalItem.ArmorValue,
		TechLevel:        internalItem.TechLevel,
		Damage:           internalItem.Damage,
		RangeMeters:      internalItem.RangeMeters,
		WeightKg:         internalItem.WeightKg,
		MagazineCapacity: internalItem.MagazineCapacity,
		AmmoType:         internalItem.AmmoType,
		Legality:         internalItem.Legality,
		CostCredits:      internalItem.CostCredits,
		Tags:             append([]string{}, internalItem.Tags...), // Create a copy
	}
}

func mapInternalEquipmentSetToGeneratorEquipmentSet(internalSet gear.EquipmentSet) EquipmentSet {
	genSet := EquipmentSet{
		Armor:   make([]GearItem, len(internalSet.Armor)),
		Weapons: make([]GearItem, len(internalSet.Weapons)),
		Tools:   make([]GearItem, len(internalSet.Tools)),
		Misc:    make([]GearItem, len(internalSet.Misc)),
	}
	for i, item := range internalSet.Armor {
		genSet.Armor[i] = mapInternalItemToGeneratorItem(item)
	}
	for i, item := range internalSet.Weapons {
		genSet.Weapons[i] = mapInternalItemToGeneratorItem(item)
	}
	for i, item := range internalSet.Tools {
		genSet.Tools[i] = mapInternalItemToGeneratorItem(item)
	}
	for i, item := range internalSet.Misc {
		genSet.Misc[i] = mapInternalItemToGeneratorItem(item)
	}
	return genSet
}

// CalculateWealthPoints determines the budget for an NPC based on SOC and CitizenCategory.
// It acts as a facade to the internal gear service.
func CalculateWealthPoints(socValue int, category CitizenCategory) int {
	npcCategory := category.toNpcCitizenCategory() // Uses existing mapper from citizen_category.go
	return gear.CalculateWealthPoints(socValue, npcCategory)
}

// GenerateEquipmentSet creates a set of equipment for an NPC based on their wealth and role.
// It acts as a facade to the internal gear service.
func GenerateEquipmentSet(wealthPoints int, role Role, npcSoc int) EquipmentSet {
	npcRole := role.toNpcRole() // Uses existing mapper from role.go
	internalEquipmentSet := gear.GenerateEquipmentSet(wealthPoints, npcRole, npcSoc)
	return mapInternalEquipmentSetToGeneratorEquipmentSet(internalEquipmentSet)
}

// TODO: Add functions to calculate wealth points based on SOC and CitizenCategory.
// TODO: Add the main gear selection algorithm.
