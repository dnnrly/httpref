package httpref

import (
	"testing"
)

func Test_Anything(t *testing.T) {
}

func TestStatusRefs_ByCode(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Statuses.ByCode(tt.name); len(got) != tt.want {
				t.Errorf("StatusRefs.ByCode() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
