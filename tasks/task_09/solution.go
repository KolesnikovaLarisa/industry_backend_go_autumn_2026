package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](
	ctx context.Context,
	workers int,
	in []T,
	fn func(context.Context, T) (R, error),
) ([]Result[R], error) {

	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	if len(in) == 0 {
		return make([]Result[R], 0), nil
	}

	results := make([]Result[R], len(in))
	workerCount := min(workers, len(in))

	jobs := make(chan int)
	var wg sync.WaitGroup

	for range workerCount {
		wg.Go(func() {
			for i := range jobs {
				value, err := fn(ctx, in[i])

				if err != nil {
					var zero R
					results[i] = Result[R]{
						Value: zero,
						Err:   err,
					}
				} else {
					results[i] = Result[R]{
						Value: value,
						Err:   nil,
					}
				}
			}
		})
	}

feed:
	for i := range in {
		if ctx.Err() != nil {
			break feed
		}

		select {
		case jobs <- i:
		case <-ctx.Done():
			break feed
		}
	}

	close(jobs)
	wg.Wait()

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	return results, nil
}
