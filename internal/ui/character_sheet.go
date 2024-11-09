package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"strings"
)

//go:generate gonstructor -type=CharacterSheet -constructorTypes=builder -output=character_sheet_auto.go

const headerIndex = -1

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
		PaddingLeft(1).
		Foreground(style.titleColor)
	valueBox := lipgloss.NewStyle().
		MarginLeft(1).
		Foreground(style.valueColor)

	nameRoleTable := table.New().
		Border(lipgloss.RoundedBorder()).
		Headers("Name", "Role").
		Rows([]string{c.fullName, fmt.Sprintf("%v (%v)", c.role, c.experience)}).
		Width(80).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch row {
			case headerIndex:
				return titleBox
			default:
				return valueBox
			}
		})

	sheetRender.WriteString(nameRoleTable.Render())
	//headers := []string{"Attribute", "Value"}
	//data := [][]string{
	//	{"Full Name", c.fullName},
	//	{"Role", c.role},
	//	{"Experience", c.experience},
	//	{"Skills", strings.Join(c.skills, ", ")},
	//}
	//
	//for characteristic, value := range c.characteristics {
	//	data = append(data, []string{characteristic.String(), fmt.Sprintf("%d", value)})
	//}
	//
	//sheet := table.New().
	//	Border(lipgloss.RoundedBorder()).
	//	Headers(headers...).
	//	Rows(data...).
	//	Width(50).
	//	StyleFunc(func(row, col int) lipgloss.Style {
	//		baseStyle := lipgloss.NewStyle().Padding(0, 1)
	//		headerStyle := baseStyle.Foreground(lipgloss.Color("252")).Bold(true)
	//		if row == 0 {
	//			return headerStyle
	//		}
	//		return baseStyle
	//	})

	return sheetRender.String()
}
