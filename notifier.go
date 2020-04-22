package iowatcher

type notifier struct {
	n chan int
}

func newNotifier() *notifier {
	return &notifier{
		n: make(chan int),
	}
}

func (n *notifier) Notifier() chan int {
	return n.n
}
