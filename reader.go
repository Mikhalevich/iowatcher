package iowatcher

import (
	"io"
)

type ReadWatcher struct {
	notifier
	origin io.Reader
}

func NewReadWatcher(r io.Reader) *ReadWatcher {
	return &ReadWatcher{
		notifier: *newNotifier(),
		origin:   r,
	}
}

func (r *ReadWatcher) Read(p []byte) (int, error) {
	n, err := r.origin.Read(p)

	if n > 0 {
		r.Notifier() <- n
	}

	if err == io.EOF {
		close(r.Notifier())
	}

	//nolint:wrapcheck
	return n, err
}

func (r *ReadWatcher) Close() error {
	close(r.Notifier())
	return nil
}
