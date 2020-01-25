package httpref

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReferences_ByName(t *testing.T) {
	Statuses = append(Statuses, Reference{Name: "501-extended"})
	type args struct {
		code string
	}
	tests := []struct {
		name string
		want int
	}{
		{name: "1", want: 5},
		{name: "40", want: 10},
		{name: "501", want: 1},
		{name: "501*", want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Statuses.ByName(tt.name); len(got) != tt.want {
				t.Errorf("References.ByName() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestReference_SummarizeContainsCorrectParts(t *testing.T) {
	r := Reference{
		Name:        "name",
		Summary:     "summary",
		Description: "description",
	}

	s := r.Summarize(100)
	assert.Contains(t, s, "name")
	assert.Contains(t, s, "summary")
}

func TestReference_SummarizeLimitsLineLength(t *testing.T) {
	r := Reference{
		Name:        "title name",
		IsTitle:     true,
		Summary:     "this is an extremely long line sfasfasdfsd werasasg asfgsdfgdsf sdfgdfs sdfg dsfg dsfg dsfg sdfg sfg a rwr sdfg sdfdffb sdfg dsg sfg dfsg sd",
		Description: "title description",
	}

	s := r.Summarize(100)
	for i, line := range strings.Split(s, "\n") {
		assert.True(t, len(line) < 100, "line %d is length %d - '%s'", i, len(line), line)
	}
}

func TestReference_DescribeLimitsLength(t *testing.T) {
	r := Headers.ByName("Headers")[0]
	description := r.Describe(100)

	assert.Contains(t, description, "HTTP")
	assert.Contains(t, description, "apply")
	for i, line := range strings.Split(description, "\n") {
		assert.True(t, len(line) < 100, "line %d is length %d - '%s'", i, len(line), line)
	}
}

func TestReferences_Titles(t *testing.T) {
	n := Statuses.Titles()
	assert.Equal(t, 5, len(n))
}
