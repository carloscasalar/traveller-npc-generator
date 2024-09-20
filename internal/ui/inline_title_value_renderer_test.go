package ui_test

import (
	"github.com/carloscasalar/traveller-npc-generator/internal/ui"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInlineTitleValueRenderer_Render_should_present_the_title_and_value_with_some_separation_in_between(t *testing.T) {
	inlineRenderer := ui.NewTitleValueRenderer("the title:", "the value")

	renderedValue := inlineRenderer.Render()

	assert.Equal(t, " the title: the value", renderedValue)
}
