package workerpool

import (
	"fmt"
	"testing"
)

func TestExecuteJobsWithWorkerPool(t *testing.T) {
	jobs := make([]Job, 0)
	for i := 0; i < 100; i++ {
		jobs = append(jobs, NewJob(i))
	}
	wp := NewWorkerPool()
	result := wp.ExecuteJobsWithWorkerPool(wp.GenerateJobs(jobs), 5)

	counts := make(map[int]int)
	for r := range result {
		counts[r.workerId]++
		fmt.Println(fmt.Sprintf("Job %d was executed by worker %d", r.jobId, r.workerId))
	}
	for key, val := range counts {
		fmt.Println(fmt.Sprintf("Worker %d executes %d jobs", key, val))
	}

}

type job struct {
	jobId int
}

func (j *job) Execute(workerId int) JobResult {
	return JobResult{workerId: workerId, jobId: j.jobId}
}

func NewJob(jobId int) Job {
	return &job{jobId: jobId}
}
