package iowatcher

import "io"

type WriteWatcher struct {
	notifier
	w io.Writer
}

func NewWriteWatcher(w io.Writer) *WriteWatcher {
	return &WriteWatcher{
		notifier: *newNotifier(),
		w:        w,
	}
}

func (ww *WriteWatcher) Write(p []byte) (int, error) {
	n, err := ww.w.Write(p)

	if n > 0 {
		ww.Notifier() <- n
	}

	return n, err
}

func (ww *WriteWatcher) Close() error {
	close(ww.Notifier())
	return nil
}
