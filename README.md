# Coding Task
##### Daimler TSS

## Installation

To be able to build this solution, run the tests/benchmarks or just run the code Go needs to be installed:
- Go (v1.16) - https://golang.org/dl/

Some `Makefile` commands need extra tools which are then listed below in the tools section.

## Usage
All use cases are added to the `Makefile`. It can be called without a command (just `make`) to list all available commands.
Following a list of the `Makefile` commands and what they do (in alphabetical order).

- `bench` runs the implemented benchmarks.
- `build` builds the example in `./cmd/sample`. It can then be executed using `./sample`.
- `depgraph` builds a png file of all Go dependencies. (`depgraph.png` added to repo, so the tools needed for this 
  do not need to be installed. If you want to run this, see `Tools` section below and install `graphViz` and `godepgraph`)
- `doc` starts an http server serving the Go documentation. Open the printed URL to jump directly to the package documentation.
- `run` uses `go run` to execute the sample in `./cmd/sample`.
- `test` executes the tests (in verbose mode).

## Answers on Task
1) Wie ist die Laufzeit Ihres Programms?
> The run time of the merge function is O(n log n).
2) Wie kann die Robustheit sichergestellt werden, vor allem auch mit Hinblick auf sehr große Eingaben?
> There are 2 factors that could break robustness: 1) response time and 2) memory.
> 1) Regarding response time: all that can be done is optimise the algorithm to be as fast as possible.
> 2) Regarding memory: the `Merge` function will allocate a maximum memory of the size of the input slice. The worst
     case is if there are no overlaps, then every interval will be added to the resulting `slice`. This could be avoided 
     by writing a merge function that merges the intervals in place (in the existing `slice` without creating a new one). 
     If we are talking the size of a list of intervals that does not event fit into memory, then we could use a streaming 
     algorithm. It would require (!) the intervals to be streamed in a sorted fashion. It would have an input stream of 
     intervals, and an output stream of intervals and remember only one interval to compare it to the next.
3) Wie verhält sich der Speicherverbrauch ihres Programs?
> The website https://afteracademy.com/blog/merge-overlapping-intervals claims that the `Merge` function is O(1) if 
> the sorting is done in place, which Go does. I don't think that is correct, as it creates a new slice to return 
> and fills it with the non-overlapping intervals. So it is limited by O(n). It can be made O(1) if the merging is also
> done in place as discussed also in the previous answer.

## Thoughts on the task itself
1) The task says to merge intervals that overlap. It is not clearly defined what happens with intervals that
   follow one another (e.g `[1,6][6,9]`) as it is also not clearly defined if start and end belong to the interval,
   or if one or both of them are excluded. For my solution I will assume that both start and end are part of the interval.
   Then the example above cosists of overlapping intervals which need to be merged.
2) It is not clearly defined if the input of the function is a `string` or `byte array` which needs to be parsed first and
   if the output should be formatted as `string`/`byte array`. In my solution I will assume we are talking about an internal function
   which gets the data as a list (`slice`) of intervals and returns a list (`slice`) of intervals (no parsing required).
3) The task says to create a "function" (in contrast to a "program"). I will primarily write tests and benchmarks 
   in order to test the function. I will also add a small hard-coded sample program that can be executed.
4) The task does not really imply that intervals will alway be valid. I will incorporate that and return an error in 
   cases where `end < start`.
5) What prerequisites do I ask for to be installed? I could go with docker and then use the Go docker image to build 
   or execute any further command. Just installing Go would remove the docker dependency, but not show off my docker skills.
   Will go with installing Go to keep the scope more focused on the task.

## Steps to the Solution
I will try to let my git history reflect my steps and additionally document them here.
1) Reflect the task itself. Think about border cases, undefined behaviour, context. In a production scenario I would
   clearify these questions with stakeholders, product managers, colleagues, etc. Here I had to make some assumptions
   which I have documented above in the section "Thoughts on the task itself".
2) Create the git repo, initialize the go module.
3) Define the module API (`Merge` function and `Interval` type).
4) Write some (failing) tests for the function.
5) Have a quick look on promising merge algorithms, and their time vs space complexity to confirm or dismiss my initial 
   idea of sorting the intervals first (avoiding O(n2)):
   - https://afteracademy.com/blog/merge-overlapping-intervals
   - https://www.csestack.org/merge-overlapping-intervals/
6) Sorting the input `slice` will change the slice, so we'd need to make a copy of it if we want a pure function. Adjust
   API to accomodate for both use cases. The non-pure function will have the advantage of more performance and less allocated
   memory.
7) first implementation getting the tests to pass
8) add benchmarks
9) add Makefile to execute tests and benchmarks; includes a list of available commands
10) Benchmarks are as expected. The pure implementation is slightly slower and allocates more memory:
```
BenchmarkMerge-12     	 5804788	       176.1 ns/op	     113 B/op	       4 allocs/op
BenchmarkMergeP-12    	 5925892	       200.8 ns/op	     160 B/op	       5 allocs/op
```
11) Cleanup and optimize; check performance against a reference implementation from SO: https://codereview.stackexchange.com/questions/259048/merge-intervalsgolang
```
BenchmarkMerge-12               	 6218521	       170.8 ns/op	     113 B/op	       4 allocs/op
BenchmarkMergeAlternative-12    	 6644235	       174.2 ns/op	     113 B/op	       4 allocs/op
```
12) Add interval validation
13) Add sample app and run/build to Makefile
14) Add depgraph
15) Complete documentation and add to Makefile
16) Answer questions of the task.
17) Write inplace algorithm reducing the memory-limitation problem.
18) Benchmarking the algorithm shows there are still allocations. Creating a memory profile (second box) shows that the 
    inplace sorting of Go is not allocation free (also discussed here https://github.com/golang/go/issues/17332). 
    The error return creates an allocation as well.
```
BenchmarkMergeInplace-12        	 8518618	       138.4 ns/op	      74 B/op	       2 allocs/op
```
```
         .          .     54:func MergeInplace(intervals []Interval) ([]Interval, error) {
         .          .     55:	if len(intervals) == 0 {
         .          .     56:		return nil, nil
         .          .     57:	}
         .          .     58:	if len(intervals) > 1 {
     150MB   530.02MB     59:		sort.Slice(intervals, func(i, j int) bool {
         .          .     60:			return intervals[i].Start < intervals[j].Start
         .          .     61:		})
         .          .     62:	}
         .          .     63:
         .          .     64:	var current int
         .          .     65:	for i, interval := range intervals {
         .          .     66:		if interval.End < interval.Start {
      19MB      101MB     67:			return nil, fmt.Errorf("invalid interval: from %d to %d", interval.Start, interval.End)
         .          .     68:		}
         .          .     69:
         .          .     70:		if i == 0 {
         .          .     71:			continue
         .          .     72:		}
```

    

# Time Table
- 7.Jun 17:40-18:30 (1. to 5.)
- 8.Jun 08:45-10:20 (6. to 13.) incl. 15min in breaks
- 8.Jun 11:00-11:30 (14. to 15.)
- 8.Jun 16:10 (16. to )

# Tools
- Language: Go
- Goland IDE (Jetbrains)
- iTerm
- git
- SourceTree
- graphViz (https://graphviz.org/)
- godepgraph (https://github.com/kisielk/godepgraph)
- godoc (https://pkg.go.dev/golang.org/x/tools/cmd/godoc)
