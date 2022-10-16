package httpref

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var (
	resultStyle, nameStyle, summaryStyle lipgloss.Style
	descriptionStyle                     *glamour.TermRenderer
	areStylesPopulated                   bool
)

type RenderStyle int64

// Summarize creates a block of text that summarizes this reference
func (r Reference) Summarize() string {
	name := nameStyle.Render(r.Name)
	summary := summaryStyle.Render(r.Summary)
	return lipgloss.JoinVertical(lipgloss.Bottom, name, summary)
}

// Describe creates a full, formated description of a reference
func (r Reference) Describe() string {
	name := nameStyle.Render(r.Name)
	summary := summaryStyle.PaddingLeft(2).Render(r.Summary)
	description, err := descriptionStyle.Render(r.Description)
	if err != nil {
		// hoping this doesn't happen as most commands here suceed without issue
		panic(err)
	}

	return lipgloss.JoinVertical(lipgloss.Bottom, name, summary, description)
}

func renderStyles(baseStyle lipgloss.Style) {
	resultStyle = baseStyle.Copy()
	nameStyle = baseStyle.Copy().
		Foreground(lipgloss.Color("86")).
		Bold(true).
		Underline(true)

	margin := uint(1)
	glamour.DarkStyleConfig.Document.Margin = &margin
	glamour.LightStyleConfig.Document.Margin = &margin

	summaryStyle = baseStyle.Copy()
	r, err := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width (default is 80)
		glamour.WithWordWrap(baseStyle.GetWidth()),
	)
	if err != nil {
		// hoping this doesn't happen as most commands here suceed without issue
		panic(err)
	}
	descriptionStyle = r
}

func PrintResultsWithStyle(results References, rootStyle lipgloss.Style) {
	// doing the population in a point where the width should not change anymore
	if !areStylesPopulated {
		areStylesPopulated = true
		renderStyles(rootStyle)
	}
	//render := getRendererByStyle(style)
	switch len(results) {
	case 0:
		res := "Filter not found any results\n"
		fmt.Fprintf(os.Stderr, res)
		os.Exit(1)
	case 1:
		fmt.Printf("%s\n", results[0].Describe())
	default:
		for _, r := range results {
			fmt.Printf("%s\n", r.Summarize())
		}
	}
}
