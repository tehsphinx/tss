package tss

import (
	"fmt"
	"sort"
)

// Interval defines a number interval. Start and End are included.
type Interval struct {
	Start, End int
}

// Merge merges the given intervals returning a list of intervals without any overlap.
// The function will return an error if one or more invalid intervals are provided.
func Merge(intervals []Interval) ([]Interval, error) {
	if len(intervals) == 0 {
		return nil, nil
	}
	if len(intervals) > 1 {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i].Start < intervals[j].Start
		})
	}

	var (
		res     []Interval
		current Interval
	)
	for i, interval := range intervals {
		if interval.End < interval.Start {
			return nil, fmt.Errorf("invalid interval: from %d to %d", interval.Start, interval.End)
		}

		if i == 0 {
			current = interval
			continue
		}
		if interval.Start <= current.End {
			current.End = max(interval.End, current.End)
			continue
		}
		res = append(res, current)
		current = interval
	}
	// add left over item
	res = append(res, current)

	return res, nil
}

// MergeAlternative merges the given intervals returning a list of intervals without any overlap.
// The function will return an error if one or more invalid intervals are provided.
func MergeAlternative(intervals []Interval) ([]Interval, error) {
	if len(intervals) > 1 {
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i].Start < intervals[j].Start
		})
	}

	var res []Interval
	for _, interval := range intervals {
		if interval.End < interval.Start {
			return nil, fmt.Errorf("invalid interval: from %d to %d", interval.Start, interval.End)
		}

		maxI := len(res) - 1
		if res == nil || interval.Start > res[maxI].End {
			res = append(res, interval)
		} else if res[maxI].End < interval.End {
			res[maxI].End = interval.End
		}
	}

	return res, nil
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

func max(i1, i2 int) int {
	if i1 < i2 {
		return i2
	}
	return i1
}
