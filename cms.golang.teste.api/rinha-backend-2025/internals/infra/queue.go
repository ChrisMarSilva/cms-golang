package infra

import (
	"errors"
	"sync"
)

type PaymentJob struct {
	CorrID string
	Cents  int64
}

type Queue struct {
	jobs chan PaymentJob
	wg   sync.WaitGroup
}

func NewQueue(capacity, workers int, handler func(PaymentJob)) *Queue {
	q := &Queue{jobs: make(chan PaymentJob, capacity)}

	for i := 0; i < workers; i++ {
		q.wg.Add(1)

		go func() {
			defer q.wg.Done()

			for job := range q.jobs {
				handler(job)
			}
		}()
	}
	return q
}

func (q *Queue) Enqueue(j PaymentJob) error {
	select {
	case q.jobs <- j:
		return nil
	default:
		return errors.New("queue full")
	}
}

func (q *Queue) Len() int {
	return len(q.jobs)
}

func (q *Queue) Shutdown() {
	close(q.jobs)
	q.wg.Wait()
}
