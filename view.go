package httpref

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var (
	resultStyle, nameStyle, summaryStyle, descriptionStyle lipgloss.Style
	passedRenderDate                                       time.Time
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
	description := descriptionStyle.PaddingTop(2).Render(r.Description)

	return lipgloss.JoinVertical(lipgloss.Bottom, name, summary, description)
}

func renderStyles(baseStyle lipgloss.Style) {
	resultStyle = baseStyle.Copy()
	nameStyle = baseStyle.Copy().
		Foreground(lipgloss.Color("86")).
		Bold(true).
		Underline(true)
	summaryStyle = baseStyle.Copy().
		PaddingLeft(2).
		Italic(true)
	descriptionStyle = baseStyle.Copy().
		Foreground(lipgloss.Color("#B2BEB5"))
}

func PrintResultsWithStyle(results References, rootStyle lipgloss.Style) {
	if passedRenderDate.IsZero() {
		passedRenderDate = time.Now()
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
