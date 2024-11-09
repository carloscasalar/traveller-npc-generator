package ui_test

import (
	"github.com/carloscasalar/traveller-npc-generator/internal/ui"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacterSheet_Render(t *testing.T) {
	// Given
	sheet := ui.NewCharacterSheetBuilder().
		FullName("Fig Kim").
		Role("trader").
		CitizenCategory("exceptional").
		Experience("elite").
		Skills([]string{"Broker-3", "Persuade-2", "Admin-2", "Advocate-1", "Electronics (Computers)-1", "Streetwise-1", "Gun Combat-0", "Diplomacy-0", "Carouse-0", "Mechanic-0", "Electronic-0", "Leadership-0"}).
		Characteristics(map[ui.Characteristic]int{
			ui.STR: 8,
			ui.DEX: 9,
			ui.END: 7,
			ui.INT: 11,
			ui.EDU: 9,
			ui.SOC: 10,
		}).
		Build()

	// When
	rendered := sheet.Render()

	// Then
	expected :=
		`╭─────────────────────────────┬────────────────────────────────────────────────╮
│ Name                        │ Role                                           │
├─────────────────────────────┼────────────────────────────────────────────────┤
│ Fig Kim                     │ Trader (Exceptional Elite)                     │
╰─────────────────────────────┴────────────────────────────────────────────────╯
╭─────────────┬────────────┬────────────┬────────────┬────────────┬────────────╮
│     STR     │    DEX     │    END     │    INT     │    EDU     │    SOC     │
├─────────────┼────────────┼────────────┼────────────┼────────────┼────────────┤
│      8      │     9      │     7      │     11     │     9      │     10     │
╰─────────────┴────────────┴────────────┴────────────┴────────────┴────────────╯
╭──────────────────────────────────────────────────────────────────────────────╮
│ Skills                                                                       │
├──────────────────────────────────────────────────────────────────────────────┤
│ Broker-3, Persuade-2, Admin-2, Advocate-1, Electronics (Computers)-1,        │
│ Streetwise-1, Gun Combat-0, Diplomacy-0, Carouse-0, Mechanic-0, Electronic-0,│
│ Leadership-0                                                                 │
╰──────────────────────────────────────────────────────────────────────────────╯`
	assert.Equal(t, expected, rendered)
}
