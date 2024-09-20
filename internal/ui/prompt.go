package ui

import "github.com/charmbracelet/lipgloss"

type PromptRenderer struct {
	style  *lipGlossStyle
	prompt string
}

func NewPromptRenderer(prompt string) *PromptRenderer {
	return &PromptRenderer{
		style:  newLipGlossDefaultStyle(),
		prompt: prompt,
	}
}

func (r PromptRenderer) Render() string {
	promptBox := lipgloss.NewStyle().
		Bold(true).
		PaddingLeft(1).
		Foreground(r.style.titleColor)
	return promptBox.Render(r.prompt)
}
