package iowatcher

import "io"

type ReadWatcher struct {
	r     io.Reader
	notifier chan int
}

func NewReadWatcher(r io.Reader) *ReadWatcher {
	return &ReadWatcher{
		r:     r,
		notifier: make(chan int),
	}
}

func (rw *ReadWatcher) Notifier() chan int {
	return rw.notifier
}

func (rw *ReadWatcher) Read(p []byte) (int, error) {
	n, err := rw.r.Read(p)

	if n > 0 {
		rw.notifier <- n
	}

	if err == io.EOF {
		close(rw.notifier)
	}

	return n, err
}