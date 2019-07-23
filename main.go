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


func dfs(ID int, job Job, list_visited []int, t_map map[int]int) []int{

    index := sort.SearchInts(list_visited, ID)
    
    if index < len(list_visited) && list_visited[index] == ID {
        // found 
    } else {

        list_visited = append(list_visited, ID)    
        for _, item := range job.childJobIDs {
            list_visited = dfs(item, jobs[t_map[ID]], list_visited, t_map)
        }
    }

    return list_visited
}


func totalExecutionTime(ID int, jobs []Job) int {
    
    var temp_visited []int
    temp_map := make(map[int]int)
    
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].ID <= jobs[j].ID
	})

    for index, value := range jobs {
        temp_map[value.ID] = index
    }
    
	temp_visited = dfs(ID, jobs[temp_map[ID]], temp_visited, temp_map)
    
	sum := 0
  
	for _, item := range temp_visited {
		sum = sum + jobs[temp_map[item]].jobTime
	}
	return sum

}

func main() {

    
	fmt.Printf("Result is: %d\n", totalExecutionTime(1, jobs)) // 120
	fmt.Printf("Result is: %d\n", totalExecutionTime(2, jobs)) // 30

}
