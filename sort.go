package tss

func quickSortInplace(vals []Interval) []Interval {
	qSort(vals, 0, len(vals)-1)
	return vals
}

func qSort(arr []Interval, start, end int) {
	if end-start < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start
	for i := start; i < end; i++ {
		if arr[i].Start < pivot.Start {
			arr[splitIndex], arr[i] = arr[i], arr[splitIndex]
			splitIndex++
		}
	}
	arr[splitIndex], arr[end] = arr[end], arr[splitIndex]

	qSort(arr, start, splitIndex-1)
	qSort(arr, splitIndex+1, end)
}

// implementationo of a slow sort algorithm to test in place sorting without allocations.
func sortInplace(vals []Interval) {
	switched := true
	for switched {
		switched = sortRun(vals)
	}
}

func sortRun(vals []Interval) bool {
	l := len(vals)
	var switched bool
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if vals[j].Start < vals[i].Start {
				vals[i], vals[j] = vals[j], vals[i]
				switched = true
			}
		}
	}
	return switched
}
