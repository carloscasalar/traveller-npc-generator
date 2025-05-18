package gear

import (
	"math/rand"
	"time"

	"github.com/carloscasalar/traveller-npc-generator/internal/npc" // For npc.Role and npc.CitizenCategory
)

// CalculateWealthPoints determines the budget for an NPC based on SOC and CitizenCategory.
func CalculateWealthPoints(socValue int, category npc.CitizenCategory) int {
	var baseWealth int
	switch {
	case socValue <= 2:
		baseWealth = 50
	case socValue <= 5:
		baseWealth = 200
	case socValue <= 8:
		baseWealth = 1000
	case socValue <= 11: // 9, A, B
		baseWealth = 5000
	case socValue <= 14: // C, D, E
		baseWealth = 20000
	default: // 15+ (F+)
		baseWealth = 100000
	}

	var categoryMultiplier float64
	switch category {
	case npc.CategoryBelowAverage:
		categoryMultiplier = 0.6
	case npc.CategoryAverage:
		categoryMultiplier = 1.0
	case npc.CategoryAboveAverage:
		categoryMultiplier = 1.5
	case npc.CategoryExceptional:
		categoryMultiplier = 2.5
	default:
		categoryMultiplier = 1.0 // Default to average if category is unknown
	}

	return int(float64(baseWealth) * categoryMultiplier)
}

// domainRoleItemTagPreferences defines preferred tags for items based on NPC role (using internal domain types).
var domainRoleItemTagPreferences = map[npc.Role][]string{
	npc.RolePilot:       {"vehicle_operation", "navigation", "spacesuit_compatible"},
	npc.RoleNavigator:   {"navigation", "computer", "sensor"},
	npc.RoleEngineer:    {"engineering", "repair", "tool", "power_systems"},
	npc.RoleSteward:     {"service", "social", "luxury"},
	npc.RoleMedic:       {"medical", "first_aid", "science", "rescue"},
	npc.RoleMarine:      {"weapon", "armor", "combat", "military", "security"},
	npc.RoleGunner:      {"weapon", "combat", "heavy_weapon", "security"},
	npc.RoleScout:       {"exploration", "survival", "sensor", "stealth"},
	npc.RoleTechnician:  {"electronics", "repair", "tool", "computer", "mechanical"},
	npc.RoleLeader:      {"communication", "social", "command"},
	npc.RoleDiplomat:    {"social", "communication", "formal_wear"},
	npc.RoleEntertainer: {"social", "performance", "art"},
	npc.RoleTrader:      {"commerce", "negotiation", "cargo"},
	npc.RoleThug:        {"weapon", "intimidation", "stealth", "makeshift"},
}

func itemHasPreferredTag(itemTags []string, preferredTags []string) bool {
	if len(preferredTags) == 0 {
		return true
	}
	for _, pTag := range preferredTags {
		for _, iTag := range itemTags {
			if iTag == pTag {
				return true
			}
		}
	}
	return false
}

func isItemPurchased(itemName string, itemType string, currentEquipment EquipmentSet) bool {
	if itemType == "misc" { // Allow duplicates for misc items
		return false
	}
	for _, armor := range currentEquipment.Armor {
		if armor.Name == itemName {
			return true
		}
	}
	for _, weapon := range currentEquipment.Weapons {
		if weapon.Name == itemName {
			return true
		}
	}
	for _, tool := range currentEquipment.Tools {
		if tool.Name == itemName {
			return true
		}
	}
	return false
}

// GenerateEquipmentSet creates a set of equipment for an NPC based on their wealth and role.
// It uses domain types for role and returns a domain EquipmentSet.
func GenerateEquipmentSet(wealthPoints int, role npc.Role, npcSoc int) EquipmentSet {
	rand.Seed(time.Now().UnixNano())

	equipment := EquipmentSet{
		Armor:   []Item{},
		Weapons: []Item{},
		Tools:   []Item{},
		Misc:    []Item{},
	}

	remainingWealth := wealthPoints
	preferredTags := domainRoleItemTagPreferences[role]

	isLawAbiding := true
	if role == npc.RoleThug {
		isLawAbiding = false
	}

	const maxIterations = 200
	const minPurchaseValue = 10

	for i := 0; i < maxIterations && remainingWealth > minPurchaseValue; i++ {
		possibleItems := make([]Item, 0)

		// Use GetMasterXxxList functions from loader.go
		allMasterItems := append(append(append(GetMasterArmorList(), GetMasterWeaponList()...), GetMasterToolList()...), GetMasterMiscList()...)

		for _, item := range allMasterItems {
			if item.CostCredits <= remainingWealth && item.CostCredits > 0 {
				if isLawAbiding && item.Legality != "Legal" && item.Legality != "" {
					continue
				}
				if !itemHasPreferredTag(item.Tags, preferredTags) {
					// Lenient: if no specific preference for role, or if item doesn't match specific but general items are OK.
					// A stricter approach might check: if len(preferredTags) > 0 && !itemHasPreferredTag(...) then continue
				}
				if isItemPurchased(item.Name, item.Type, equipment) {
					continue
				}
				possibleItems = append(possibleItems, item)
			}
		}

		if len(possibleItems) == 0 {
			break
		}

		selectedItem := possibleItems[rand.Intn(len(possibleItems))]

		switch selectedItem.Type {
		case "armor":
			equipment.Armor = append(equipment.Armor, selectedItem)
		case "weapon":
			equipment.Weapons = append(equipment.Weapons, selectedItem)
		case "tool":
			equipment.Tools = append(equipment.Tools, selectedItem)
		case "misc":
			equipment.Misc = append(equipment.Misc, selectedItem)
		}
		remainingWealth -= selectedItem.CostCredits
	}

	return equipment
}
