package pingobot

import (
	"sync"
	"time"
)

type Pool struct {
	worker     *worker
	workersNum int

	jobs    chan string
	results chan Result

	wg      *sync.WaitGroup
	stopped bool
}

func New(workersNum int, timeout time.Duration, result chan Result) *Pool {
	return &Pool{
		worker:     newWorker(timeout),
		workersNum: workersNum,

		jobs:    make(chan string),
		results: result,

		wg: new(sync.WaitGroup),
	}
}

func (p *Pool) Start() {
	for i := 1; i <= p.workersNum; i++ {
		go func(i int) {
			for job := range p.jobs {
				p.results <- p.worker.handle(job)
				p.wg.Done()
			}
		}(i)
	}
}

func (p *Pool) Push(job string) {
	if p.stopped {
		return
	}
	p.jobs <- job
	p.wg.Add(1)
}

func (p *Pool) Stop() {
	p.stopped = true
	close(p.jobs)
	p.wg.Wait()
}
