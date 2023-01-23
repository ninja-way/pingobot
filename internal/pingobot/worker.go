package pingobot

import (
	"net/http"
	"time"
)

type worker struct {
	client *http.Client
}

func newWorker(timeout time.Duration) *worker {
	return &worker{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (w worker) handle(job string) Result {
	start := time.Now()

	resp, err := w.client.Get(job)
	if err != nil {
		return Result{Error: err}
	}

	return Result{
		URL:          job,
		StatusCode:   resp.StatusCode,
		ResponseTime: time.Since(start),
	}
}
