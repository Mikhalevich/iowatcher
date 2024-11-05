package iowatcher

type BytesProcessedCallback func(bytesProcessed int)

type notifier struct {
	callback BytesProcessedCallback
}

func newNotifier(callback BytesProcessedCallback) *notifier {
	return &notifier{
		callback: callback,
	}
}

func (n *notifier) notify(bytesProcessed int) {
	n.callback(bytesProcessed)
}
