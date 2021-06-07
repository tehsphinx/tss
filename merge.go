package tss

// Interval defines a number interval. Start and End are included.
type Interval struct {
	Start, End int
}

// Merge merges the given intervals returning a list of intervals without any overlap.
// This is a pure function (not changing the input slice). The function will return an error if
// one or more invalid intervals are provided.
func Merge(intervals []Interval) ([]Interval, error) {
	return intervals, nil
}
