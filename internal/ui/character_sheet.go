package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
	citizenCategory string
	experience      string
	skills          []string
	characteristics map[Characteristic]int
}

func (c *CharacterSheet) Render() string {
	sheetRender := new(strings.Builder)

	titleBox, valueBox := c.getTableStyles()
	nameRoleTable := c.buildNameRoleTable(valueBox, titleBox)
	characteristicsTable := c.buildCharacteristicsTable(titleBox, valueBox)
	skillsTable := c.buildSillsTable(titleBox, valueBox)

	sheetRender.WriteString(nameRoleTable.Render())
	sheetRender.WriteString("\n")
	sheetRender.WriteString(characteristicsTable.Render())
	sheetRender.WriteString("\n")
	sheetRender.WriteString(skillsTable.Render())

	return sheetRender.String()
}

func (c *CharacterSheet) buildSillsTable(titleBox lipgloss.Style, valueBox lipgloss.Style) *table.Table {
	skillsTable := table.New().
		Border(lipgloss.RoundedBorder()).
		Headers("Skills").
		Rows([]string{strings.Join(c.skills, ", ")}).
		Width(tableWidth).
		StyleFunc(func(row, _ int) lipgloss.Style {
			switch row {
			case headerIndex:
				return titleBox.PaddingLeft(1)
			default:
				return valueBox.
					PaddingLeft(1).
					Height(3)
			}
		})
	return skillsTable
}

func (c *CharacterSheet) buildCharacteristicsTable(titleBox lipgloss.Style, valueBox lipgloss.Style) *table.Table {
	characteristicsTable := table.New().
		Border(lipgloss.RoundedBorder()).
		Headers(toStringList(CharacteristicValues())...).
		Rows(c.characteristicsValues()).
		Width(tableWidth).
		StyleFunc(func(row, _ int) lipgloss.Style {
			switch row {
			case headerIndex:
				return titleBox.Align(lipgloss.Center)
			default:
				return valueBox.Align(lipgloss.Center)
			}
		})
	return characteristicsTable
}

func (c *CharacterSheet) buildNameRoleTable(valueBox lipgloss.Style, titleBox lipgloss.Style) *table.Table {
	nameRoleTable := table.New().
		Border(lipgloss.RoundedBorder()).
		Headers("Name", "Role").
		Rows([]string{c.fullName, c.roleDescriptionWithStyle(valueBox)}).
		Width(tableWidth).
		StyleFunc(func(row, _ int) lipgloss.Style {
			switch row {
			case headerIndex:
				return titleBox.PaddingLeft(1)
			default:
				return valueBox.PaddingLeft(1)
			}
		})
	return nameRoleTable
}

func (c *CharacterSheet) getTableStyles() (titleBoxStyle lipgloss.Style, valueBoxStyle lipgloss.Style) {
	style := newLipGlossDefaultStyle()
	titleBoxStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(style.titleColor)
	valueBoxStyle = lipgloss.NewStyle().
		Foreground(style.valueColor)
	return
}

func (c *CharacterSheet) roleDescriptionWithStyle(baseStyle lipgloss.Style) string {
	roleBox := baseStyle
	categoryExperienceBox := baseStyle.Italic(true).MarginLeft(1)
	categoryBox := baseStyle.Italic(true).Faint(true)
	experienceBox := baseStyle.Italic(true).Bold(true).MarginLeft(1)
	titleCaser := cases.Title(language.English)

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		roleBox.Render(titleCaser.String(c.role)),
		categoryExperienceBox.Render(
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				"(",
				categoryBox.Render(titleCaser.String(c.citizenCategory)),
				experienceBox.Render(titleCaser.String(c.experience)),
				baseStyle.Render(")"),
			),
		),
	)
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
