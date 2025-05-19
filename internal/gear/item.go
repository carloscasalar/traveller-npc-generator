package gear

// Item represents a generic piece of equipment in the domain.
type Item struct {
	Name             string   `yaml:"name"`
	Type             string   `yaml:"type"` // e.g., "armor", "weapon", "tool", "misc"
	WeaponType       string   `yaml:"weapon_type,omitempty"`
	ArmorValue       int      `yaml:"armor_value,omitempty"`
	TechLevel        int      `yaml:"tech_level"`
	Damage           string   `yaml:"damage,omitempty"`
	RangeMeters      int      `yaml:"range_meters,omitempty"`
	MagazineCapacity int      `yaml:"magazine_capacity,omitempty"`
	Legality         string   `yaml:"legality"`
	CostCredits      int      `yaml:"cost_credits"`
	Tags             []string `yaml:"tags"`
}

// EquipmentSet holds all gear for an NPC in the domain.
type EquipmentSet struct {
	Armor   []Item
	Weapons []Item
	Tools   []Item
	Misc    []Item
}
