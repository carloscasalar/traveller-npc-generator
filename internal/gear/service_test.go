package gear

import (
	"testing"

	"github.com/carloscasalar/traveller-npc-generator/internal/npc" // For npc.CitizenCategory
	"github.com/stretchr/testify/assert"                            // Added testify/assert
	"github.com/stretchr/testify/require"
)

func TestCalculateWealthPoints(t *testing.T) {
	testCases := []struct {
		name           string
		socValue       int
		category       npc.CitizenCategory
		expectedWealth int
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
			assert.Equal(t, tc.expectedWealth, actualWealth)
		})
	}
}

// TODO: Add more detailed tests for GenerateEquipmentSet (role preference, legality etc.)
// TODO: Add tests for helper functions if they become more complex or are exported (itemHasPreferredTag, isItemPurchased)

func TestGenerateEquipmentSet_ZeroWealth(t *testing.T) {
	require.NoError(t, Init())

	equipment := GenerateEquipmentSet(0, npc.RolePilot, 7, npc.CategoryAverage)

	assert.Empty(t, equipment.Armor)
	assert.Empty(t, equipment.Weapons)
	assert.Empty(t, equipment.Tools)
	assert.Empty(t, equipment.Misc)
}

func TestGenerateEquipmentSet_MarineWithHighWealth(t *testing.T) {
	require.NoError(t, Init())

	equipment := GenerateEquipmentSet(1000000, npc.RoleMarine, 8, npc.CategoryAverage)

	assert.NotEmpty(t, equipment.Armor)
	assert.NotEmpty(t, equipment.Weapons)
	assert.NotEmpty(t, equipment.Tools)
	assert.NotEmpty(t, equipment.Misc)
}

func TestGenerateEquipmentSet_MarineEquipmentMatchesRole(t *testing.T) {
	require.NoError(t, Init())

	equipment := GenerateEquipmentSet(1000000, npc.RoleMarine, 8, npc.CategoryAverage)

	marineTags := domainRoleItemTagPreferences[npc.RoleMarine]
	hasMatchingTag := false

	for _, item := range equipment.Armor {
		for _, tag := range item.Tags {
			for _, marineTag := range marineTags {
				if tag == marineTag {
					hasMatchingTag = true
					break
				}
			}
		}
	}

	assert.True(t, hasMatchingTag, "Marine equipment should have at least one matching role tag")
}

func TestGenerateEquipmentSet_MarineEquipmentMatchesQuality(t *testing.T) {
	require.NoError(t, Init())

	equipment := GenerateEquipmentSet(1000000, npc.RoleMarine, 8, npc.CategoryAverage)

	qualityTags := GetQualityTags(8, npc.CategoryAverage)
	hasMatchingTag := false

	for _, item := range equipment.Armor {
		for _, tag := range item.Tags {
			for _, qualityTag := range qualityTags {
				if tag == qualityTag {
					hasMatchingTag = true
					break
				}
			}
		}
	}

	assert.True(t, hasMatchingTag, "Marine equipment should have at least one matching quality tag")
}

func TestGenerateEquipmentSet_MarineEquipmentRespectsLegality(t *testing.T) {
	require.NoError(t, Init())

	equipment := GenerateEquipmentSet(1000000, npc.RoleMarine, 8, npc.CategoryAverage)

	for _, item := range equipment.Armor {
		assert.True(t, canAccessLegality(npc.RoleMarine, item.Legality),
			"Marine equipment should respect legality restrictions")
	}
	for _, item := range equipment.Weapons {
		assert.True(t, canAccessLegality(npc.RoleMarine, item.Legality),
			"Marine equipment should respect legality restrictions")
	}
}
