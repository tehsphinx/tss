package tss

import (
	"testing"

	"github.com/matryer/is"
)

var mergeTests = []struct {
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
	{
		name: "invalid interval",
		input: []Interval{
			{Start: -22, End: -23},
		},
		want:    nil,
		wantErr: true,
	},
}

func TestMerge(t *testing.T) {
	for _, tt := range mergeTests {
		t.Run(tt.name, func(t *testing.T) {
			asrt := is.New(t)

			got, err := Merge(tt.input)
			// check if error is as expected
			asrt.Equal(err != nil, tt.wantErr)
			// compare expected result
			asrt.Equal(got, tt.want)
		})
	}
}

func TestMergeInplace(t *testing.T) {
	for _, tt := range mergeTests {
		t.Run(tt.name, func(t *testing.T) {
			asrt := is.New(t)

			got, err := MergeInplace(tt.input)
			// check if error is as expected
			asrt.Equal(err != nil, tt.wantErr)
			// compare expected result
			asrt.Equal(got, tt.want)
		})
	}
}

func TestMergeAlternative(t *testing.T) {
	for _, tt := range mergeTests {
		t.Run(tt.name, func(t *testing.T) {
			asrt := is.New(t)

			got, err := MergeAlternative(tt.input)
			// check if error is as expected
			asrt.Equal(err != nil, tt.wantErr)
			// compare expected result
			asrt.Equal(got, tt.want)
		})
	}
}

func TestMergeP(t *testing.T) {
	for _, tt := range mergeTests {
		t.Run(tt.name, func(t *testing.T) {
			asrt := is.New(t)

			// make a copy of input to be able to compare it against the original for changes by reference
			input := make([]Interval, len(tt.input))
			copy(input, tt.input)
			if tt.input == nil {
				input = nil
			}

			got, err := MergeP(input)
			// check if error is as expected
			asrt.Equal(err != nil, tt.wantErr)
			// compare expected result
			asrt.Equal(got, tt.want)
			// check if input was changed
			asrt.Equal(input, tt.input)
		})
	}
}

func BenchmarkMerge(b *testing.B) {
	l := len(mergeTests)
	for i := 0; i < b.N; i++ {
		tt := mergeTests[i%l]
		_, _ = Merge(tt.input)
	}
}

func BenchmarkMergeInplace(b *testing.B) {
	l := len(mergeTests)
	for i := 0; i < b.N; i++ {
		tt := mergeTests[i%l]
		_, _ = MergeInplace(tt.input)
	}
}

func BenchmarkMergeAlternative(b *testing.B) {
	l := len(mergeTests)
	for i := 0; i < b.N; i++ {
		tt := mergeTests[i%l]
		_, _ = MergeAlternative(tt.input)
	}
}

func BenchmarkMergeP(b *testing.B) {
	l := len(mergeTests)
	for i := 0; i < b.N; i++ {
		tt := mergeTests[i%l]
		_, _ = MergeP(tt.input)
	}
}
