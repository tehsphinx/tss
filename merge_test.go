package tss

import (
	"testing"

	"github.com/matryer/is"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name    string
		input   []Interval
		want    []Interval
		wantErr bool
	}{
		{
			name:  "no intervals",
			input: nil,
			want:  nil,
		},
		{
			name: "no overlapping intervals",
			input: []Interval{
				{Start: -22, End: -1},
				{Start: 0, End: 3},
				{Start: 234, End: 5234},
			},
			want: []Interval{
				{Start: -22, End: -1},
				{Start: 0, End: 3},
				{Start: 234, End: 5234},
			},
		},
		{
			name: "two overlapping intervals",
			input: []Interval{
				{Start: -22, End: -1},
				{Start: -5, End: 3},
				{Start: 234, End: 5234},
			},
			want: []Interval{
				{Start: -22, End: 3},
				{Start: 234, End: 5234},
			},
		},
		{
			name: "included interval",
			input: []Interval{
				{Start: -22, End: -1},
				{Start: 523, End: 2352},
				{Start: 234, End: 5234},
			},
			want: []Interval{
				{Start: -22, End: -1},
				{Start: 234, End: 5234},
			},
		},
		{
			name: "overlapping and included",
			input: []Interval{
				{Start: -22, End: -1},
				{Start: -20, End: -10},
				{Start: -5, End: 33},
				{Start: -3, End: 29},
			},
			want: []Interval{
				{Start: -22, End: 33},
			},
		},
		{
			name: "adjacent intervals",
			input: []Interval{
				{Start: -20, End: -1},
				{Start: -1, End: 1},
				{Start: 1, End: 15},
			},
			want: []Interval{
				{Start: -20, End: 15},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asrt := is.New(t)

			got, err := Merge(tt.input)
			asrt.Equal(err != nil, tt.wantErr)
			asrt.Equal(got, tt.want)
		})
	}
}
