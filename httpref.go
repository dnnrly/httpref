package httpref

import (
	"fmt"
	"strings"

	para "github.com/dnnrly/paragraphical"
	"github.com/muesli/termenv"
)

// Reference is a single reference item
type Reference struct {
	Name        string
	IsTitle     bool
	Summary     string
	Description string
}

// References is a collection of Reference objects
type References []Reference

// Summarize creates a block of text that summarizes this reference
func (r Reference) Summarize(width int) string {
	name := termenv.String(r.Name)
	summary := termenv.String(r.Summary).Italic()
	return para.Format(
		width,
		fmt.Sprintf("%s\n  %s", name.Bold().Underline(), summary),
	)
}

// Describe creates a full, formated description of a reference
func (r Reference) Describe(width int) string {
	text := fmt.Sprintf(
		"%s\n  %s\n\n%s",
		termenv.String(r.Name).Bold().Underline(),
		termenv.String(r.Summary).Italic(),
		r.Description,
	)

	return para.Format(width, text)
}

// ByName finds all of the Reference with a matching Name field
func (r References) ByName(code string) References {
	results := References{}

	wildcard := strings.HasSuffix(code, "*")
	code = strings.ReplaceAll(code, "*", "")

	for _, v := range r {
		if !wildcard && v.Name == code {
			return References{v}
		}

		if strings.HasPrefix(v.Name, code) {
			results = append(results, v)
		}
	}

	return results
}

// InRange attempts to find a numeric in a range in the Name field
func (r References) InRange(code string) References {
	for _, v := range r {
		if strings.Contains(v.Name, "-") {
			parts := strings.Split(v.Name, "-")
			if code >= parts[0] && code <= parts[len(parts)-1] {
				return References{v}
			}
		}
	}

	return References{}
}

// Titles gives all of the Reference objects marked as a title
func (r References) Titles() References {
	results := References{}

	for _, v := range r {
		if v.IsTitle {
			results = append(results, v)
		}
	}

	return results
}

func (r References) Search(term string) References {
	results := References{}

	for _, v := range r {
		if strings.Contains(strings.ToLower(v.Name), term) ||
			strings.Contains(strings.ToLower(v.Summary), term) ||
			strings.Contains(strings.ToLower(v.Description), term) {
			results = append(results, v)
		}
	}

	return results
}
