package ui

//go:generate gonstructor -type=CharacterSheet -constructorTypes=builder -output=character_sheet_auto.go

// CharacterSheet Is a struct that can render the character sheet in a text format
type CharacterSheet struct {
	fullName        string
	role            string
	experience      string
	skills          []string
	characteristics map[Characteristic]int
}

func (c *CharacterSheet) Render() string {
	return "Character Sheet"
}
