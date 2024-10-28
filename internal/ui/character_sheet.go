package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"strings"
)

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
	headers := []string{"Attribute", "Value"}
	data := [][]string{
		{"Full Name", c.fullName},
		{"Role", c.role},
		{"Experience", c.experience},
		{"Skills", strings.Join(c.skills, ", ")},
	}

	for characteristic, value := range c.characteristics {
		data = append(data, []string{characteristic.String(), fmt.Sprintf("%d", value)})
	}

	sheet := table.New().
		Border(lipgloss.NormalBorder()).
		Headers(headers...).
		Rows(data...).
		Width(50).
		StyleFunc(func(row, col int) lipgloss.Style {
			baseStyle := lipgloss.NewStyle().Padding(0, 1)
			headerStyle := baseStyle.Foreground(lipgloss.Color("252")).Bold(true)
			if row == 0 {
				return headerStyle
			}
			return baseStyle
		})

	return sheet.String()
}
