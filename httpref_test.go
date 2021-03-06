package httpref

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReferences_ByName(t *testing.T) {
	Statuses = append(Statuses, Reference{Name: "501-extended"})
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

func TestReferences_InRange(t *testing.T) {
	tests := []struct {
		having string
		expect string
	}{
		{having: "19150", expect: "10000-20000"},
		{having: "16406", expect: "10000-20000"},
		{having: "5988", expect: "5988-5989"},
		{having: "5989", expect: "5988-5989"},
	}

	for _, tt := range tests {
		t.Run(tt.having, func(t *testing.T) {
			got := RegisteredPorts.InRange(tt.having)
			if len(got) != 1 {
				t.Errorf("References.InRange() = %v, want %v", len(got), 1)
			}
			if got[0].Name != tt.expect {
				t.Errorf("References.InRange()[0] = %v, want %v", got[0].Name, tt.expect)
			}
		})
	}

	t.Run("Invalid port  should not return anything", func(t *testing.T) {
		got := RegisteredPorts.InRange("70000") // Invalid port, should not return anything
		if len(got) != 0 {
			t.Errorf("References.InRange() = %v, want %v", len(got), 0)
		}
	})
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

func TestPortsConsistencyValidation(t *testing.T) {
	ports := append(WellKnownPorts[1:], RegisteredPorts[1:]...)
	var validRange = regexp.MustCompile(`^\d+(-\d+)?$`)
	for _, port := range ports {
		if !validRange.MatchString(port.Name) {
			t.Errorf("Invalid port format: %v", port.Name)
		}
	}
}

func TestReferences_Search(t *testing.T) {
	n := Statuses.Search("auth")
	assert.Equal(t, 8, len(n))
}
