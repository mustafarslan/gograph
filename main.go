package main

import (
	"fmt"
	"sort"
)

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
