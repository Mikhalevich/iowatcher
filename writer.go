package iowatcher

import "io"

type WriteWatcher struct {
	*notifier
	origin io.Writer
}

func NewWriteWatcher(w io.Writer, callback BytesProcessedCallback) *WriteWatcher {
	return &WriteWatcher{
		notifier: newNotifier(callback),
		origin:   w,
	}
}

func (w *WriteWatcher) Write(p []byte) (int, error) {
	n, err := w.origin.Write(p)

	if n > 0 {
		w.notify(n)
	}

	//nolint:wrapcheck
	return n, err
}
