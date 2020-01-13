package httpref

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Summary(t *testing.T) {
	r := Reference{
		Name:        "name",
		Summary:     "summary",
		Description: "description",
	}

	s := r.Summarize()
	assert.Equal(t, "                          name summary", s)

	r = Reference{
		Name:        "title name",
		IsTitle:     true,
		Summary:     "title summary",
		Description: "title description",
	}

	s = r.Summarize()
	assert.Equal(t, "                    title name title summary", s)
}

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
