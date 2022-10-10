package httpref

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			PaddingTop(1)
		//			PaddingLeft(4)
	resultStyle = baseStyle.Copy()
)

const width = 100

type RenderStyle int64

const (
	Summary    RenderStyle = 0
	SearchTerm             = 1
	Titles                 = 2
	ByName                 = 3
)

func getRendererByStyle(style RenderStyle) func(string) string {
	switch style {
	case Summary:
		return resultStyle.Render
	default:
		return func(str string) string {
			return fmt.Sprintf("%s\n", str)
		}
	}
}
func PrintResultsWithStyle(results References, style RenderStyle) {
	render := getRendererByStyle(style)
	switch len(results) {
	case 0:
		res := render("Filter not found any results\n")
		fmt.Fprintf(os.Stderr, res)
		os.Exit(1)
	case 1:
		fmt.Printf("%s\n", render(results[0].Describe(width)))
	default:
		for _, r := range results {
			fmt.Printf("%s\n", render(r.Summarize(width)))
		}
	}
}
