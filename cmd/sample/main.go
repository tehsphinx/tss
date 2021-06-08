// This package implements a sample program merging a predefined set of intervals.
package main

import (
	"fmt"
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/tehsphinx/tss"
)

func main() {
	intervals := []tss.Interval{
		{Start: 110, End: 130},
		{Start: -3, End: 29},
		{Start: 100, End: 103},
		{Start: -20, End: -10},
		{Start: -3, End: 29},
		{Start: -22, End: -1},
		{Start: 110, End: 120},
		{Start: -20, End: -10},
		{Start: -5, End: 33},
		{Start: 105, End: 110},
	}

	fmt.Println("Input Intervals:")
	for _, interval := range intervals {
		fmt.Printf("\t%+v\n", aurora.Cyan(interval))
	}

	merged, err := tss.Merge(intervals)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Output Intervals:")
	for _, interval := range merged {
		fmt.Printf("\t%+v\n", aurora.Green(interval))
	}
}
