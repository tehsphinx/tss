# Coding Task
##### Daimler TSS

## Thoughts on the task itself
1) The task says to merge intervals that overlap. It is not clearly defined what happens with intervals that
   follow one another (e.g `[1,6][6,9]`) as it is also not clearly defined if the start and the end belong to the interval
   or if one or both of them are excluded. For my solution I will assume that both start and end are part of the interval.
   Then the example above cosists of overlapping intervals which need to be merged.
2) It is not clearly defined if the input of the function is a `string` or `byte array` which needs to be parsed first and
   if the output should be formatted as `string`/`byte array`. In my solution I will assume we are talking about an internal function
   which gets the data as a list (`slice`) of intervals and returns a list (`slice`) of intervals (no parsing required).
3) The task says to create a "function" (in contrast to a "program"). I will primarily write tests and benchmarks 
   in order to test the function. I will also add a small hard-coded sample program that can be executed.
4) The task does not really imply that intervals will alway be valid. I will incorporate that and return an error in 
   cases where `end < start`.

## Steps to the Solution
I will try to let my git history reflect my steps and additionally document them here.
1) Reflect the task itself. Think about border cases, undefined behaviour, context. In a production scenario I would
   clearify these questions with stake holders, product managers, colleagues, etc. Here I had to make some assumptions
   which I have documented above in the section "Thoughts on the task itself".
2) Create the git repo, initialize the go module.
3) Define the module API (`Merge` function and `Interval` type).
4) Write some (failing) tests for the function.
5) Have a quick look on promising merge algorithms, and their time vs space complexity to confirm or dismiss my initial 
   idea of sorting the intervals first:
   - https://afteracademy.com/blog/merge-overlapping-intervals
   - https://www.csestack.org/merge-overlapping-intervals/

# Time Table
- 7.Jun 17:30-18:30
