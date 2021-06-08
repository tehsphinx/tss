package tss

import "sort"

// Interval defines a number interval. Start and End are included.
type Interval struct {
	Start, End int
}

// Merge merges the given intervals returning a list of intervals without any overlap.
// The function will return an error if one or more invalid intervals are provided.
func Merge(intervals []Interval) ([]Interval, error) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var (
		res     []Interval
		current Interval
	)
	for i, interval := range intervals {
		if i == 0 {
			current = interval
			continue
		}
		if interval.Start <= current.End {
			current.End = max(interval.End, current.End)
			if i == len(intervals)-1 {
				res = append(res, current)
			}
			continue
		}
		res = append(res, current)
		current = interval
		if i == len(intervals)-1 {
			res = append(res, current)
		}
	}

	return res, nil
}

func max(i1, i2 int) int {
	if i1 < i2 {
		return i2
	}
	return i1
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
