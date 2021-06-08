package tss

import (
	"errors"
	"fmt"
	"sort"
)

// ErrorInvalidInterval is returned if an invalid interval is provided. Invalid in the sense that end < start.
var ErrorInvalidInterval = errors.New("invalid interval: start is larger than end")

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
			// overlap
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

// MergeInplace merges the given intervals returning a list of intervals without any overlap.
// The function will return an error if one or more invalid intervals are provided.
// This uses quick sort for sorting.
func MergeInplace(intervals []Interval) ([]Interval, error) {
	if len(intervals) == 0 {
		return nil, nil
	}
	if len(intervals) > 1 {
		quickSortInplace(intervals)
	}

	var current int
	for i, interval := range intervals {
		if interval.End < interval.Start {
			return nil, ErrorInvalidInterval
		}

		if i == 0 {
			continue
		}
		if interval.Start <= intervals[current].End {
			// overlap
			intervals[current].End = max(interval.End, intervals[current].End)
			continue
		}
		current++
		intervals[current].Start = interval.Start
		intervals[current].End = interval.End
	}

	return intervals[:current+1], nil
}

// MergeInplaceBasicSort merges the given intervals returning a list of intervals without any overlap.
// The function will return an error if one or more invalid intervals are provided.
// Using a rudementary sorting algorithm.
func MergeInplaceBasicSort(intervals []Interval) ([]Interval, error) {
	if len(intervals) == 0 {
		return nil, nil
	}
	if len(intervals) > 1 {
		sortInplace(intervals)
	}

	var current int
	for i, interval := range intervals {
		if interval.End < interval.Start {
			return nil, ErrorInvalidInterval
		}

		if i == 0 {
			continue
		}
		if interval.Start <= intervals[current].End {
			// overlap
			intervals[current].End = max(interval.End, intervals[current].End)
			continue
		}
		current++
		intervals[current].Start = interval.Start
		intervals[current].End = interval.End
	}

	return intervals[:current+1], nil
}

// MergeStream merges a stream of intervals, streaming a list of intervals back without any overlap.
// This function expects an already sorted (by start) and sanitized stream of valid intervals.
// The output channel is closed if the input was closed and all remaining items have been streamed.
func MergeStream(in <-chan Interval, out chan<- Interval) {
	defer close(out)

	// get first element and return if there is none (channel closed)
	current, ok := <-in
	if !ok {
		return
	}
	for interval := range in {
		if interval.Start <= current.End {
			// overlap
			current.End = max(interval.End, current.End)
			continue
		}

		out <- current
		current = interval
	}
	out <- current
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
			// overlap
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
