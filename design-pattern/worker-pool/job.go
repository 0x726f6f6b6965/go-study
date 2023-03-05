package workerpool

type Job interface {
	Execute(workerId int) JobResult
}

type JobResult struct {
	workerId int
	jobId    int
}
