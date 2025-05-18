package gear

import (
	"fmt"

	"github.com/carloscasalar/traveller-npc-generator/assets"
	"gopkg.in/yaml.v3"
)

var (
	masterArmorList  []Item
	masterWeaponList []Item
	masterToolList   []Item
	masterMiscList   []Item
)

// Init loads all master gear lists from embedded data provided by the assets package.
// This should be called once during application startup.
func Init() error {
	var err error
	err = yaml.Unmarshal(assets.GetEmbeddedArmorData(), &masterArmorList)
	if err != nil {
		return fmt.Errorf("failed to unmarshal embedded armor data: %w", err)
	}
	err = yaml.Unmarshal(assets.GetEmbeddedWeaponsData(), &masterWeaponList)
	if err != nil {
		return fmt.Errorf("failed to unmarshal embedded weapon data: %w", err)
	}
	err = yaml.Unmarshal(assets.GetEmbeddedToolsData(), &masterToolList)
	if err != nil {
		return fmt.Errorf("failed to unmarshal embedded tool data: %w", err)
	}
	err = yaml.Unmarshal(assets.GetEmbeddedMiscData(), &masterMiscList)
	if err != nil {
		return fmt.Errorf("failed to unmarshal embedded misc items data: %w", err)
	}
	return nil
}

// GetMasterArmorList returns a copy of the master armor list.
func GetMasterArmorList() []Item {
	listCopy := make([]Item, len(masterArmorList))
	copy(listCopy, masterArmorList)
	return listCopy
}

// GetMasterWeaponList returns a copy of the master weapon list.
func GetMasterWeaponList() []Item {
	listCopy := make([]Item, len(masterWeaponList))
	copy(listCopy, masterWeaponList)
	return listCopy
}

// GetMasterToolList returns a copy of the master tool list.
func GetMasterToolList() []Item {
	listCopy := make([]Item, len(masterToolList))
	copy(listCopy, masterToolList)
	return listCopy
}

// GetMasterMiscList returns a copy of the master misc list.
func GetMasterMiscList() []Item {
	listCopy := make([]Item, len(masterMiscList))
	copy(listCopy, masterMiscList)
	return listCopy
}
