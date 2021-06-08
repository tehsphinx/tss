package tss

// Interval defines a number interval. Start and End are included.
type Interval struct {
	Start, End int
}

// Merge merges the given intervals returning a list of intervals without any overlap.
// The function will return an error if one or more invalid intervals are provided.
func Merge(intervals []Interval) ([]Interval, error) {
	return intervals, nil
}

// MergeP merges the given intervals returning a list of intervals without any overlap.
// This is a pure function (not changing the input slice). The function will return an error if
// one or more invalid intervals are provided.
func MergeP(intervals []Interval) ([]Interval, error) {
	intervals = copyIntervals(intervals)
	return Merge(intervals)
}

func copyIntervals(intervals []Interval) []Interval {
	c := make([]Interval, len(intervals))
	copy(c, intervals)
	return c
}
