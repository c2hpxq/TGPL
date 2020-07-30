// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 276.

// Package memo provides a concurrency-safe memoization a function of
// a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a Mutex.
package memo

import "sync"

// Func is the type of the function to memoize.
type Func func(string, chan struct{}) (interface{}, error, bool)
type RawFunc func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//!+
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

func New(f RawFunc) *Memo {
	ff := func (s string, c chan struct{}) (interface{}, error, bool) {
		fDone := make(chan result)
		var r result
		go func() {
			r.value, r.err = f(s)
			close(fDone)
		}()

		select {
		case <-fDone:
			return r.value, r.err, false
		case <-c:
			return nil, nil, true
		}
	}
	return &Memo{f: ff, cache: make(map[string]*entry)}
}

type Memo struct {
	f     Func
	mu    sync.Mutex // guards cache
	cache map[string]*entry
}

func (memo *Memo) Get(key string, done chan struct{}) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		var cancelled bool
		e.res.value, e.res.err, cancelled = memo.f(key, done)

		if cancelled {
			memo.cache[key] = nil
		}

		close(e.ready) // broadcast ready condition
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
		select {
			case <-e.ready: // wait for ready condition
			case <-done:
				return nil, nil
		}
	}
	return e.res.value, e.res.err
}

//!-
