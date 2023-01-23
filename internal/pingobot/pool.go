package pingobot

import (
	"time"
)

type Pool struct {
	worker     *worker
	workersNum int

	jobs    chan string
	results chan Result
}

func New(workersNum int, timeout time.Duration, result chan Result) *Pool {
	return &Pool{
		worker:     newWorker(timeout),
		workersNum: workersNum,

		jobs:    make(chan string),
		results: result,
	}
}

func (p *Pool) Start() {
	for i := 1; i <= p.workersNum; i++ {
		go func() {
			for job := range p.jobs {
				p.results <- p.worker.handle(job)
			}
		}()
	}
}

func (p *Pool) Push(job string) {
	p.jobs <- job
}
