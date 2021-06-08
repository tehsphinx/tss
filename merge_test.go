package tss

import (
	"sort"
	"testing"

	"github.com/matryer/is"
)

type test struct {
	name    string
	input   []Interval
	want    []Interval
	wantErr bool
}

func getTests() []test {
	return []test{
		{
			name:  "no intervals",
			input: nil,
			want:  nil,
		},
		{
			name: "no overlapping intervals",
			input: []Interval{
				{Start: 0, End: 3},
				{Start: -22, End: -1},
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
				{Start: 234, End: 5234},
				{Start: -22, End: -1},
				{Start: -5, End: 3},
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
				{Start: -3, End: 29},
				{Start: -22, End: -1},
				{Start: -5, End: 33},
				{Start: -20, End: -10},
			},
			want: []Interval{
				{Start: -22, End: 33},
			},
		},
		{
			name: "adjacent intervals",
			input: []Interval{
				{Start: -1, End: 1},
				{Start: 1, End: 15},
				{Start: -20, End: -1},
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
}

func TestMerge(t *testing.T) {
	for _, tt := range getTests() {
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
	for _, tt := range getTests() {
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

func TestMergeStream(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				// MergeStream does not handle invalid intervals, so we skip tests expecting an error.
				t.Skip()
			}

			asrt := is.New(t)

			chSend := make(chan Interval)
			chRes := make(chan Interval)

			go MergeStream(chSend, chRes)
			go func() {
				input := tt.input
				sort.Slice(input, func(i, j int) bool {
					return input[i].Start < input[j].Start
				})

				for _, interval := range input {
					chSend <- interval
				}
				close(chSend)
			}()

			var got []Interval
			for interval := range chRes {
				got = append(got, interval)
			}

			// compare expected result
			asrt.Equal(got, tt.want)
		})
	}
}

func TestMergeAlternative(t *testing.T) {
	for _, tt := range getTests() {
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
	for _, tt := range getTests() {
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
	mergeTests := getTests()
	l := len(mergeTests)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tt := mergeTests[i%l]
		_, _ = Merge(tt.input)
	}
}

func BenchmarkMergeInplace(b *testing.B) {
	mergeTests := getTests()
	l := len(mergeTests)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tt := mergeTests[i%l]
		_, _ = MergeInplace(tt.input)
	}
}

func BenchmarkMergeAlternative(b *testing.B) {
	mergeTests := getTests()
	l := len(mergeTests)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tt := mergeTests[i%l]
		_, _ = MergeAlternative(tt.input)
	}
}

func BenchmarkMergeP(b *testing.B) {
	mergeTests := getTests()
	l := len(mergeTests)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tt := mergeTests[i%l]
		_, _ = MergeP(tt.input)
	}
}
