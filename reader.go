package iowatcher

import "io"

type ReadWatcher struct {
	notifier
	r     io.Reader
}

func NewReadWatcher(r io.Reader) *ReadWatcher {
	return &ReadWatcher{
		notifier: *newNotifier(),
		r:     r,
	}
}

func (rw *ReadWatcher) Read(p []byte) (int, error) {
	n, err := rw.r.Read(p)

	if n > 0 {
		rw.Notifier() <- n
	}

	if err == io.EOF {
		close(rw.Notifier())
	}

	return n, err
}
