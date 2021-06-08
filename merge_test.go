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

func TestMergeInplaceBasicSort(t *testing.T) {
	for _, tt := range getTests() {
		t.Run(tt.name, func(t *testing.T) {
			asrt := is.New(t)

			got, err := MergeInplaceBasicSort(tt.input)
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
	data := getBenchData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Merge(data)
	}
}

func BenchmarkMergeInplace(b *testing.B) {
	data := getBenchData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MergeInplace(data)
	}
}

func BenchmarkMergeInplaceBasicSort(b *testing.B) {
	data := getBenchData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MergeInplaceBasicSort(data)
	}
}

func BenchmarkMergeAlternative(b *testing.B) {
	data := getBenchData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MergeAlternative(data)
	}
}

func BenchmarkMergeP(b *testing.B) {
	data := getBenchData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = MergeP(data)
	}
}

func getBenchData() []Interval {
	return []Interval{
		{Start: 110, End: 130},
		{Start: 100, End: 103},
		{Start: -3, End: 29},
		{Start: -34513, End: -35513},
		{Start: -20, End: -10},
		{Start: -22, End: -1},
		{Start: -3492, End: -35129},
		{Start: -34723, End: -35784},
		{Start: -34923, End: -35932},
		{Start: -3, End: 29},
		{Start: -20, End: -10},
		{Start: 110, End: 120},
		{Start: 500, End: 523},
		{Start: -5, End: 33},
		{Start: 513, End: 513},
		{Start: -34643, End: -35734},
		{Start: 512, End: 521},
		{Start: 92, End: 129},
		{Start: 723, End: 784},
		{Start: 123, End: 237},
		{Start: 923, End: 932},
		{Start: 1105, End: 1110},
		{Start: 1512, End: 1521},
		{Start: 1500, End: 1523},
		{Start: 1513, End: 1513},
		{Start: 1723, End: 1784},
		{Start: 1643, End: 1734},
		{Start: -31923, End: -31932},
		{Start: -31123, End: -31237},
		{Start: -34500, End: -35523},
		{Start: -34105, End: -35110},
		{Start: -34512, End: -35521},
		{Start: -341105, End: -351110},
		{Start: 105, End: 110},
		{Start: -34123, End: -35237},
		{Start: -341500, End: -351523},
		{Start: -341513, End: -351513},
		{Start: -341512, End: -351521},
		{Start: -341643, End: -351734},
		{Start: 192, End: 1129},
		{Start: -341723, End: -351784},
		{Start: 1923, End: 1932},
		{Start: 4105, End: 5110},
		{Start: 1123, End: 1237},
		{Start: 643, End: 734},
		{Start: 41723, End: 51784},
		{Start: 41923, End: 51932},
		{Start: 4192, End: 51129},
		{Start: 41123, End: 51237},
		{Start: -3500, End: -3523},
		{Start: -3105, End: -3110},
		{Start: -3512, End: -3521},
		{Start: -3643, End: -3734},
		{Start: -3513, End: -3513},
		{Start: -3723, End: -3784},
		{Start: -3923, End: -3932},
		{Start: -392, End: -3129},
		{Start: -3123, End: -3237},
		{Start: 4513, End: 5513},
		{Start: -31105, End: -31110},
		{Start: 4643, End: 5734},
		{Start: 492, End: 5129},
		{Start: 4723, End: 5784},
		{Start: 4923, End: 5932},
		{Start: -31500, End: -31523},
		{Start: 4123, End: 5237},
		{Start: -31513, End: -31513},
		{Start: -31512, End: -31521},
		{Start: -31643, End: -31734},
		{Start: -3192, End: -31129},
		{Start: -31723, End: -31784},
		{Start: -34192, End: -351129},
		{Start: -341123, End: -351237},
		{Start: -341923, End: -351932},
		{Start: 4500, End: 5523},
		{Start: 41105, End: 51110},
		{Start: 4512, End: 5521},
		{Start: 41500, End: 51523},
		{Start: 41643, End: 51734},
		{Start: 41513, End: 51513},
		{Start: 41512, End: 51521},
	}
}
