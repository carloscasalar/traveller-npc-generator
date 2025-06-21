package gear

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitAndGetMasterLists(t *testing.T) {
	// Call Init to load data from embedded files
	err := Init()
	assert.NoError(t, err, "Init() should not return an error")

	// Test Armor List
	armorList := GetMasterArmorList()
	assert.NotEmpty(t, armorList, "Expected master armor list to be populated")

	if len(armorList) > 0 { // Guarding access for the specific item check
		expectedFirstName := "Flak Jacket"
		assert.Equal(t, expectedFirstName, armorList[0].Name, "First armor item name mismatch")

		// Check if it's a copy
		originalName := armorList[0].Name
		armorList[0].Name = "TestChangeShouldNotAffectMaster" // Modify the copy
		// Fetch master list again or access directly if test setup allows (masterArmorList is package-level)
		assert.NotEqual(t, masterArmorList[0].Name, armorList[0].Name, "GetMasterArmorList did not return a copy, modification affected master list")
		assert.Equal(t, originalName, masterArmorList[0].Name, "Master list was unexpectedly changed")
		// No need to restore armorList[0].Name as it's a distinct copy
		// Restore masterArmorList[0].Name if it was modified for some reason (it shouldn't be by armorList modification)
		// For clarity, if masterArmorList[0].Name was indeed originalName, this re-asserts it.
		// The critical part is that masterArmorList[0].Name is NOT "TestChangeShouldNotAffectMaster"
	}

	// Test Weapon List
	weaponList := GetMasterWeaponList()
	assert.NotEmpty(t, weaponList, "Expected master weapon list to be populated")
	// TODO: Add a specific item check if desired, e.g., assert.Equal(t, "ExpectedName", weaponList[0].Name)

	// Test Tool List
	toolList := GetMasterToolList()
	assert.NotEmpty(t, toolList, "Expected master tool list to be populated")
	// TODO: Add a specific item check if desired

	// Test Misc List
	miscList := GetMasterMiscList()
	assert.NotEmpty(t, miscList, "Expected master misc list to be populated")
	// TODO: Add a specific item check if desired
}
