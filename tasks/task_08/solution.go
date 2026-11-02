package main

import (
	"sync"
	"time"
)

type Clock interface{ Now() time.Time }
type Limiter struct {
	mu     sync.Mutex
	clock  Clock
	rate   float64
	burst  int
	tokens float64
	last   time.Time
}

func NewLimiter(clock Clock, ratePerSec float64, burst int) *Limiter {
	l := Limiter{
		clock:  clock,
		rate:   ratePerSec,
		burst:  burst,
		tokens: float64(burst),
	}

	if clock != nil {
		l.last = clock.Now()
	}

	return &l
}
func (l *Limiter) AllowN(n int) bool {
	if l.burst <= 0 || l.clock == nil || n <= 0 || n > l.burst {
		return false
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	now := l.clock.Now()

	if !now.Before(l.last) {
		elap := now.Sub(l.last).Seconds()
		if l.rate > 0 {
			l.tokens = min(float64(l.burst), l.tokens+l.rate*elap)
		}
		l.last = now
	}

	if l.tokens-float64(n) >= 0 {
		l.tokens -= float64(n)
		return true
	}

	return false
}
