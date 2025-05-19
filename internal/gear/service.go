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
	npc.RoleMarine:      {"weapon", "armor", "combat", "military", "security", "communication", "survival"},
	npc.RoleGunner:      {"weapon", "combat", "heavy_weapon", "security", "communication"},
	npc.RoleScout:       {"exploration", "survival", "sensor", "stealth", "communication"},
	npc.RoleTechnician:  {"electronics", "repair", "tool", "computer", "mechanical", "communication"},
	npc.RoleLeader:      {"communication", "social", "command", "survival"},
	npc.RoleDiplomat:    {"social", "communication", "formal_wear"},
	npc.RoleEntertainer: {"social", "performance", "art"},
	npc.RoleTrader:      {"commerce", "negotiation", "cargo", "communication"},
	npc.RoleThug:        {"weapon", "intimidation", "stealth", "makeshift", "survival"},
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

// TagCategory represents the type of tag and its source
type TagCategory int

const (
	TagCategoryPurpose TagCategory = iota // Tags from roles indicating needs/purposes
	TagCategoryQuality                    // Tags from SOC/citizen category indicating quality/access
	TagCategoryBase                       // Base item tags
)

// Tag represents a single tag with its category
type Tag struct {
	Name     string
	Category TagCategory
}

// TagSet represents a collection of tags from different sources
type TagSet struct {
	PurposeTags []string // Tags from role indicating needs/purposes
	QualityTags []string // Tags from SOC/citizen category indicating quality/access
	BaseTags    []string // Base item tags
}

// GetRolePurposeTags returns the purpose tags for a given role
func GetRolePurposeTags(role npc.Role) []string {
	return domainRoleItemTagPreferences[role]
}

// GetQualityTags returns quality/access tags based on SOC and citizen category
func GetQualityTags(socValue int, category npc.CitizenCategory) []string {
	var qualityTags []string

	// Add quality tags based on SOC
	switch {
	case socValue <= 2:
		qualityTags = append(qualityTags, "basic", "low_quality")
	case socValue <= 5:
		qualityTags = append(qualityTags, "standard", "common")
	case socValue <= 8:
		qualityTags = append(qualityTags, "good", "reliable")
	case socValue <= 11:
		qualityTags = append(qualityTags, "high_quality", "premium")
	case socValue <= 14:
		qualityTags = append(qualityTags, "luxury", "exclusive")
	default:
		qualityTags = append(qualityTags, "elite", "exceptional")
	}

	// Add access tags based on citizen category
	switch category {
	case npc.CategoryBelowAverage:
		qualityTags = append(qualityTags, "restricted_access")
	case npc.CategoryAverage:
		qualityTags = append(qualityTags, "standard_access")
	case npc.CategoryAboveAverage:
		qualityTags = append(qualityTags, "preferred_access")
	case npc.CategoryExceptional:
		qualityTags = append(qualityTags, "privileged_access")
	}

	return qualityTags
}

// itemMatchesTagSet checks if an item matches the given tag set
func itemMatchesTagSet(itemTags []string, tagSet TagSet) bool {
	// Check purpose tags (at least one match required)
	purposeMatch := false
	if len(tagSet.PurposeTags) == 0 {
		purposeMatch = true
	} else {
		for _, pTag := range tagSet.PurposeTags {
			for _, iTag := range itemTags {
				if pTag == iTag {
					purposeMatch = true
					break
				}
			}
			if purposeMatch {
				break
			}
		}
	}
	if !purposeMatch {
		return false
	}

	// Only require a quality tag match if the item has any quality/access tags
	itemHasQuality := false
	for _, iTag := range itemTags {
		if iTag == "basic" || iTag == "low_quality" || iTag == "standard" || iTag == "common" || iTag == "good" || iTag == "reliable" || iTag == "high_quality" || iTag == "premium" || iTag == "luxury" || iTag == "exclusive" || iTag == "elite" || iTag == "exceptional" || iTag == "restricted_access" || iTag == "standard_access" || iTag == "preferred_access" || iTag == "privileged_access" {
			itemHasQuality = true
			break
		}
	}

	if itemHasQuality {
		qualityMatch := false
		if len(tagSet.QualityTags) == 0 {
			qualityMatch = true
		} else {
			for _, qTag := range tagSet.QualityTags {
				for _, iTag := range itemTags {
					if qTag == iTag {
						qualityMatch = true
						break
					}
				}
				if qualityMatch {
					break
				}
			}
		}
		if !qualityMatch {
			return false
		}
	}

	return true
}

// canAccessLegality determines if a role can access an item with a given legality
func canAccessLegality(role npc.Role, legality string) bool {
	switch legality {
	case "Military":
		return role == npc.RoleMarine || role == npc.RoleGunner || role == npc.RoleThug
	case "Restricted":
		return role == npc.RoleMarine || role == npc.RoleGunner || role == npc.RoleThug
	case "Legal", "":
		return true
	default:
		return false
	}
}

// GenerateEquipmentSet creates a set of equipment for an NPC based on their wealth, role, and category.
func GenerateEquipmentSet(wealthPoints int, role npc.Role, npcSoc int, category npc.CitizenCategory) EquipmentSet {
	rand.Seed(time.Now().UnixNano())

	equipment := EquipmentSet{
		Armor:   []Item{},
		Weapons: []Item{},
		Tools:   []Item{},
		Misc:    []Item{},
	}

	remainingWealth := wealthPoints
	tagSet := TagSet{
		PurposeTags: GetRolePurposeTags(role),
		QualityTags: GetQualityTags(npcSoc, category),
	}

	const maxIterations = 200
	const minPurchaseValue = 10

	for i := 0; i < maxIterations && remainingWealth > minPurchaseValue; i++ {
		possibleItems := make([]Item, 0)

		// Use GetMasterXxxList functions from loader.go
		allMasterItems := append(append(append(GetMasterArmorList(), GetMasterWeaponList()...), GetMasterToolList()...), GetMasterMiscList()...)

		for _, item := range allMasterItems {
			if item.CostCredits <= remainingWealth && item.CostCredits > 0 {
				if !canAccessLegality(role, item.Legality) {
					continue
				}
				if !itemMatchesTagSet(item.Tags, tagSet) {
					continue
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
