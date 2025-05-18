package generator

import (
	"testing"

	"github.com/carloscasalar/traveller-npc-generator/internal/gear" // For internal gear types
	"github.com/stretchr/testify/assert"
	// "github.com/carloscasalar/traveller-npc-generator/internal/npc" // Might be needed if creating complex test data
)

func TestMapInternalItemToGeneratorItem(t *testing.T) {
	internalItem := gear.Item{
		Name:        "Laser Pistol",
		Type:        "weapon",
		WeaponType:  "pistol",
		TechLevel:   10,
		CostCredits: 500,
		Tags:        []string{"energy", "personal"},
	}

	expectedGenItem := GearItem{
		Name:        "Laser Pistol",
		Type:        "weapon",
		WeaponType:  "pistol",
		TechLevel:   10,
		CostCredits: 500,
		Tags:        []string{"energy", "personal"},
	}

	actualGenItem := mapInternalItemToGeneratorItem(internalItem)
	assert.Equal(t, expectedGenItem, actualGenItem, "Item mapping failed")

	// Test that tags are copied, not referenced
	if len(internalItem.Tags) > 0 && len(actualGenItem.Tags) > 0 {
		originalTagValue := internalItem.Tags[0]
		internalItem.Tags[0] = "changed_in_source"
		assert.NotEqual(t, internalItem.Tags[0], actualGenItem.Tags[0],
			"Tags slice was not copied; modification to source internalItem.Tags affected mapped actualGenItem.Tags")
		assert.Equal(t, originalTagValue, actualGenItem.Tags[0], "Original tag value in mapped item changed unexpectedly")
		internalItem.Tags[0] = originalTagValue // Restore source for consistency if other tests use it
	}
}

func TestMapInternalEquipmentSetToGeneratorEquipmentSet(t *testing.T) {
	internalSet := gear.EquipmentSet{
		Armor: []gear.Item{
			{Name: "Combat Armor", Type: "armor", CostCredits: 5000, Tags: []string{"military"}},
		},
		Weapons: []gear.Item{
			{Name: "Laser Rifle", Type: "weapon", CostCredits: 1200, Tags: []string{"military", "energy"}},
			{Name: "Knife", Type: "weapon", CostCredits: 25, Tags: []string{"melee"}},
		},
		Tools: []gear.Item{},
		Misc: []gear.Item{
			{Name: "Medkit", Type: "misc", CostCredits: 100, Tags: []string{"medical"}},
		},
	}

	expectedGenSet := EquipmentSet{
		Armor: []GearItem{
			{Name: "Combat Armor", Type: "armor", CostCredits: 5000, Tags: []string{"military"}},
		},
		Weapons: []GearItem{
			{Name: "Laser Rifle", Type: "weapon", CostCredits: 1200, Tags: []string{"military", "energy"}},
			{Name: "Knife", Type: "weapon", CostCredits: 25, Tags: []string{"melee"}},
		},
		Tools: []GearItem{},
		Misc: []GearItem{
			{Name: "Medkit", Type: "misc", CostCredits: 100, Tags: []string{"medical"}},
		},
	}

	actualGenSet := mapInternalEquipmentSetToGeneratorEquipmentSet(internalSet)
	assert.Equal(t, expectedGenSet, actualGenSet, "EquipmentSet mapping failed")

	// Spot check tag copying for one item to ensure deep copy behavior for items within the set
	if len(internalSet.Weapons) > 0 && len(actualGenSet.Weapons) > 0 &&
		len(internalSet.Weapons[0].Tags) > 0 && len(actualGenSet.Weapons[0].Tags) > 0 {
		originalTagValue := internalSet.Weapons[0].Tags[0]
		internalSet.Weapons[0].Tags[0] = "ThisShouldNotPropagateToMappedSet"
		assert.NotEqual(t, internalSet.Weapons[0].Tags[0], actualGenSet.Weapons[0].Tags[0],
			"Tags slice within an item in EquipmentSet was not copied deeply")
		assert.Equal(t, originalTagValue, actualGenSet.Weapons[0].Tags[0], "Original tag value in mapped item changed unexpectedly")
		internalSet.Weapons[0].Tags[0] = originalTagValue // Restore source
	}
}
