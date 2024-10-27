package httpref

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var (
	nameStyle, summaryStyle lipgloss.Style
	descriptionStyle        *glamour.TermRenderer
	descriptionBorderStyle  lipgloss.Style
)

type RenderStyle int64

// Summarize creates a block of text that summarizes this reference
func (r Reference) Summarize(style lipgloss.Style) string {
	name := nameStyle.Inherit(style).Render(r.Name)
	summary := summaryStyle.Inherit(style).Render(r.Summary)
	statusBar := renderStatusBar(r.Name, r.Summary)
	return lipgloss.JoinVertical(lipgloss.Bottom, name, summary, statusBar)
}

// Describe creates a full, formated description of a reference
func (r Reference) Describe(style lipgloss.Style) string {
	statusBar := renderStatusBar(r.Name, r.Summary)
	descriptionStyle, err := updateTermRendered(style)
	if err != nil {
		// hoping this doesn't happen as most commands here suceed without issue
		panic(err)
	}
	description, err := descriptionStyle.Render(r.Description)
	if err != nil {
		panic(err)
	}
	descriptionWithBorder := descriptionBorderStyle.Render(description)

	return lipgloss.JoinVertical(lipgloss.Bottom, statusBar, descriptionWithBorder)
}

func init() {
	renderStyles()
}

func renderStyles() {
	resultStyle := lipgloss.NewStyle()
	descriptionForeColorDarkTheme := "120"
	descriptionForeColorLightTheme := "202"

	descriptionBorderStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "5"}).
		Padding(0, 1, 0, 1)

	margin := uint(1)
	glamour.DarkStyleConfig.Document.Margin = &margin
	glamour.LightStyleConfig.Document.Margin = &margin
	glamour.DarkStyleConfig.Text.Color = &descriptionForeColorDarkTheme
	glamour.LightStyleConfig.Text.Color = &descriptionForeColorLightTheme

	r, err := updateTermRendered(resultStyle)
	if err != nil {
		// hoping this doesn't happen as most commands here suceed without issue
		panic(err)
	}
	descriptionStyle = r
}

func updateTermRendered(style lipgloss.Style) (*glamour.TermRenderer, error) {
	width := style.GetWidth()
	if width == 0 {
		width = 80
	}
	r, err := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width (default is 80)
		glamour.WithWordWrap(width),
	)
	return r, err
}

func PrintResultsWithStyle(results References, rootStyle lipgloss.Style) {
	switch len(results) {
	case 0:
		res := "Filter not found any results\n"
		fmt.Fprintf(os.Stderr, res)
		os.Exit(1)
	case 1:
		fmt.Printf("%s\n", results[0].Describe(rootStyle))
	default:
		for _, r := range results {
			fmt.Printf("%s\n", r.Summarize(rootStyle))
		}
	}
}

func renderStatusBar(name, summary string) string {
	// styled using adaptive color to set one of the two colors based on the current theme
	statusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF5F87")).
		Background(lipgloss.Color("#353533")).
		Bold(true).
		Padding(0, 1)

	statusText := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "20", Dark: "178"}).
		Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"}).
		Padding(0, 2)

	statusKey := statusStyle.Render(name)
	statusVal := statusText.Render(summary)

	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusKey,
		statusVal,
	)

	return statusStyle.Width(101).Render(bar)
}
