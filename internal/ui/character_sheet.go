package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"strings"
)

//go:generate gonstructor -type=CharacterSheet -constructorTypes=builder -output=character_sheet_auto.go

const (
	headerIndex = -1
	tableWidth  = 80
)

// CharacterSheet Is a struct that can render the character sheet in a text format
type CharacterSheet struct {
	fullName        string
	role            string
	experience      string
	skills          []string
	characteristics map[Characteristic]int
}

func (c *CharacterSheet) Render() string {
	sheetRender := new(strings.Builder)

	style := newLipGlossDefaultStyle()
	titleBox := lipgloss.NewStyle().
		Bold(true).
		Foreground(style.titleColor)
	valueBox := lipgloss.NewStyle().
		Foreground(style.valueColor)

	nameRoleTable := table.New().
		Border(lipgloss.RoundedBorder()).
		Headers("Name", "Role").
		Rows([]string{c.fullName, fmt.Sprintf("%v (%v)", c.role, c.experience)}).
		Width(tableWidth).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch row {
			case headerIndex:
				return titleBox.PaddingLeft(1)
			default:
				return valueBox.PaddingLeft(1)
			}
		})
	sheetRender.WriteString(nameRoleTable.Render())
	sheetRender.WriteString("\n")

	characteristicsTable := table.New().
		Border(lipgloss.RoundedBorder()).
		Headers(toStringList(CharacteristicValues())...).
		Rows(c.characteristicsValues()).
		Width(tableWidth).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch row {
			case headerIndex:
				return titleBox.Align(lipgloss.Center)
			default:
				return valueBox.Align(lipgloss.Center)
			}
		})

	sheetRender.WriteString(characteristicsTable.Render())
	sheetRender.WriteString("\n")

	skillsTable := table.New().
		Border(lipgloss.RoundedBorder()).
		Headers("Skills").
		Rows([]string{strings.Join(c.skills, ", ")}).
		Width(tableWidth).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch row {
			case headerIndex:
				return titleBox.PaddingLeft(1)
			default:
				return valueBox.
					PaddingLeft(1).
					Height(3)
			}
		})
	sheetRender.WriteString(skillsTable.Render())

	return sheetRender.String()
}

func (c *CharacterSheet) characteristicsValues() []string {
	values := make([]string, len(CharacteristicValues()))
	for i, characteristic := range CharacteristicValues() {
		values[i] = fmt.Sprintf("%v", c.characteristics[characteristic])
	}
	return values
}

func toStringList[T interface{ String() string }](values []T) []string {
	stringValues := make([]string, len(values))
	for i, v := range values {
		stringValues[i] = v.String()
	}
	return stringValues
}
