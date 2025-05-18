package gear

import (
	"testing"

	"github.com/carloscasalar/traveller-npc-generator/internal/npc" // For npc.CitizenCategory
	"github.com/stretchr/testify/assert"                            // Added testify/assert
)

func TestCalculateWealthPoints(t *testing.T) {
	testCases := []struct {
		name           string
		socValue       int
		category       npc.CitizenCategory
		expectedWealth int // Simplified to single expected value
	}{
		{"Low SOC, Below Average", 2, npc.CategoryBelowAverage, 30},
		{"Low SOC, Average", 2, npc.CategoryAverage, 50},
		{"Mid SOC, Above Average", 7, npc.CategoryAboveAverage, 1500},
		{"High SOC, Exceptional", 12, npc.CategoryExceptional, 50000},
		{"SOC 5, Average", 5, npc.CategoryAverage, 200},
		{"SOC 8, Exceptional", 8, npc.CategoryExceptional, 2500},
		{"SOC 11, Below Average", 11, npc.CategoryBelowAverage, 3000},
		{"SOC 14, Average", 14, npc.CategoryAverage, 20000},
		{"SOC 15, Exceptional", 15, npc.CategoryExceptional, 250000},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualWealth := CalculateWealthPoints(tc.socValue, tc.category)
			assert.Equal(t, tc.expectedWealth, actualWealth,
				"For SOC %d and Category %v", tc.socValue, tc.category)
		})
	}
}

// TODO: Add more detailed tests for GenerateEquipmentSet (role preference, legality etc.)
// TODO: Add tests for helper functions if they become more complex or are exported (itemHasPreferredTag, isItemPurchased)

func TestGenerateEquipmentSet_ZeroWealth(t *testing.T) {
	// Ensure master lists are loaded for this test run if not already
	if err := Init(); err != nil && len(GetMasterArmorList()) == 0 { // Simple check if Init likely hasn't run
		assert.FailNow(t, "Init() failed for test setup: %v", err) // Use FailNow for setup issues
	}

	equipment := GenerateEquipmentSet(0, npc.RolePilot, 7) // Role and SOC are arbitrary for zero wealth
	assert.Empty(t, equipment.Armor, "Armor should be empty with zero wealth")
	assert.Empty(t, equipment.Weapons, "Weapons should be empty with zero wealth")
	assert.Empty(t, equipment.Tools, "Tools should be empty with zero wealth")
	assert.Empty(t, equipment.Misc, "Misc items should be empty with zero wealth")
}

func TestGenerateEquipmentSet_SufficientWealth(t *testing.T) {
	if err := Init(); err != nil && len(GetMasterArmorList()) == 0 {
		assert.FailNow(t, "Init() failed for test setup: %v", err)
	}

	equipment := GenerateEquipmentSet(1000000, npc.RoleMarine, 8)
	// Check if at least one category of equipment is not empty
	hasSomeEquipment := len(equipment.Armor) > 0 ||
		len(equipment.Weapons) > 0 ||
		len(equipment.Tools) > 0 ||
		len(equipment.Misc) > 0

	assert.True(t, hasSomeEquipment, "Expected some equipment with high wealth for a Marine, but got none. Equipment: %+v", equipment)
	// A more specific test would check for a weapon or armor for a Marine.
}
