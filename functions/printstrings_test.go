package functions

import (
	"reflect"
	"testing"
)

func TestProcessString(t *testing.T) {

	tests := []struct {
		name     string
		args     []int
		filename string
		want     []string
	}{
		{"#",
			[]int{29, 30},
			"standard.txt",
			[]string{`  _  _      _| || |_ `},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessString(tt.args, tt.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintStrings(t *testing.T) {

	tests := []struct {
		name string
		str []string
	
	}{
		{"#",
		[]string{"   _  _      _| || |_   |_  __  _|   _| || |_   |_  __  _|    |_||_|                           "},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintStrings(tt.str)
		})
	}
}
