package assets

import (
	_ "embed" // Required for go:embed
)

//go:embed armor.yml
var EmbeddedArmorData []byte

//go:embed weapons.yml
var EmbeddedWeaponsData []byte

//go:embed tools.yml
var EmbeddedToolsData []byte

//go:embed misc_items.yml
var EmbeddedMiscData []byte

// GetEmbeddedArmorData returns the raw embedded armor data.
func GetEmbeddedArmorData() []byte {
	return EmbeddedArmorData
}

// GetEmbeddedWeaponsData returns the raw embedded weapons data.
func GetEmbeddedWeaponsData() []byte {
	return EmbeddedWeaponsData
}

// GetEmbeddedToolsData returns the raw embedded tools data.
func GetEmbeddedToolsData() []byte {
	return EmbeddedToolsData
}

// GetEmbeddedMiscData returns the raw embedded misc data.
func GetEmbeddedMiscData() []byte {
	return EmbeddedMiscData
}
