package errgroup

import (
	"sync"
)

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
//
// This Group implementation differs from sync/errgroup in that
// instead of store only first encountered error it stores all encountered errors.
//
// Can only used once, you may only call Wait once!
type Group struct {
	wg   sync.WaitGroup
	mu   sync.Mutex
	errs []error
}

// Go calls passed function in goroutine.
// Every caught error will be returned in []error by Wait.
func (g *Group) Go(f func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			g.mu.Lock()
			defer g.mu.Unlock()

			g.errs = append(g.errs, err)
		}
	}()
}

// Wait blocks until all spawned goroutines in Go method have completed,
// it returnes nil if no error encountered or []error if any encountered.
func (g *Group) Wait() []error {
	g.wg.Wait()
	g.mu.Lock()
	defer g.mu.Unlock()

	if len(g.errs) == 0 {
		return nil
	}
	return g.errs
}
