package main

import (
	"fmt"
	"sort"
)

// Given a job struct definition below write a function that takes 2 inputs:
// * ID of a job to run
// * a slice/array/list of known job descriptions
// and return a total execution time of job with ID=ID and all of its children
// recursively.
//
// We can assume that there are no cycles in dependency graph and job
// descriptions in the slice/array/list are unique i.e. there are no duplicates.
//
// Add a couple of tests to prove your solution.
// Use whatever language you feel most comfortable with.
//
// For the exemplar 'jobs' slice below:
// * given ID=4 the function should return 60 - just a single job's duration
//   without any dependencies
// * given ID=2 the function should return 30 - duration of jobs 2 and 3
// * given ID=1 the function should return 120 - duration of jobs 1, 2, 3 and 4
// * given ID=5 the function should return 200 - duration of jobs 1, 2, 3 and 4, 5,6
type Job struct {
	ID          int
	jobTime     int
	childJobIDs []int
}

var jobs = []Job{
	Job{1, 30, []int{2, 4}},
	Job{2, 10, []int{3}},
	Job{4, 60, []int{}},
	Job{3, 20, []int{}},
	Job{5, 60, []int{1, 6}},
	Job{6, 20, []int{}},
	Job{8, 40, []int{6, 2}},
	Job{7, 10, []int{5}},
	Job{9, 50, []int{2, 4, 6}},
}

var visited []int

func dfs(ID int, job Job) {

	index := sort.SearchInts(visited, ID)
	if index < len(visited) && visited[index] == ID {
		// found
	} else {
		visited = append(visited, ID)
		for _, item := range job.childJobIDs {
			pos := sort.Search(len(jobs), func(i int) bool { return jobs[i].ID >= item })
			dfs(item, jobs[pos])
		}
	}
}

func totalExecutionTime(ID int, jobs []Job) int {

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].ID <= jobs[j].ID
	})

	pos := sort.Search(len(jobs), func(i int) bool { return jobs[i].ID >= ID })
	dfs(ID, jobs[pos])

	sum := 0

	for _, item := range visited {
		pos := sort.Search(len(jobs), func(i int) bool { return jobs[i].ID >= item })
		sum = sum + jobs[pos].jobTime
	}
	return sum

}

func main() {

	fmt.Printf("Result is: %d", totalExecutionTime(1, jobs))

}
