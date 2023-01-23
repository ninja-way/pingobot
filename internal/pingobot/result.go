package pingobot

import (
	"fmt"
	"time"
)

type Result struct {
	URL          string
	StatusCode   int
	ResponseTime time.Duration
	Error        error
}

// String return result data in convenient format
func (r Result) String() string {
	if r.Error != nil {
		return fmt.Sprintf("[ERROR] %s : %s", r.URL, r.Error.Error())
	}
	return fmt.Sprintf("[OK] %s : STATUS %d : TIME %s", r.URL, r.StatusCode, r.ResponseTime)
}
