package iowatcher

import (
	"io"
)

type ReadWatcher struct {
	*notifier
	origin io.Reader
}

func NewReadWatcher(r io.Reader, callback BytesProcessedCallback) *ReadWatcher {
	return &ReadWatcher{
		notifier: newNotifier(callback),
		origin:   r,
	}
}

func (r *ReadWatcher) Read(p []byte) (int, error) {
	n, err := r.origin.Read(p)

	if n > 0 {
		r.notify(n)
	}

	//nolint:wrapcheck
	return n, err
}
