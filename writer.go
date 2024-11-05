package iowatcher

import "io"

type WriteWatcher struct {
	notifier
	origin io.Writer
}

func NewWriteWatcher(w io.Writer) *WriteWatcher {
	return &WriteWatcher{
		notifier: *newNotifier(),
		origin:   w,
	}
}

func (w *WriteWatcher) Write(p []byte) (int, error) {
	n, err := w.origin.Write(p)

	if n > 0 {
		w.Notifier() <- n
	}

	//nolint:wrapcheck
	return n, err
}

func (w *WriteWatcher) Close() error {
	close(w.Notifier())
	return nil
}
