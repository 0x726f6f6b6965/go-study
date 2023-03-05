package workerpool

import (
	"sync"
)

type WorkerPool interface {
	GenerateJobs([]Job) chan Job
	ExecuteJobsWithWorkerPool(jobs chan Job, maxWorkers int) chan JobResult
}

type workerPool struct{}

func NewWorkerPool() WorkerPool {
	return &workerPool{}
}

func (wp *workerPool) GenerateJobs(in []Job) chan Job {
	jobs := make(chan Job)
	go func() {
		defer close(jobs)
		for _, job := range in {
			jobs <- job
		}
	}()
	return jobs
}

func (wp *workerPool) ExecuteJobsWithWorkerPool(jobs chan Job, maxWorkers int) chan JobResult {
	result := make(chan JobResult)
	go func() {
		defer close(result)
		var wg sync.WaitGroup
		defer wg.Wait()
		for i := 0; i < maxWorkers; i++ {
			wg.Add(1)
			go func(workerId int) {
				defer wg.Done()
				for job := range jobs {
					result <- job.Execute(workerId)
				}
			}(i)
		}
	}()
	return result
}
