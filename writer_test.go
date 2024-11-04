package iowatcher

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrite4KBytes(t *testing.T) {
	var (
		buf     = bytes.Buffer{}
		watcher = NewWriteWatcher(&buf)
	)

	go func() {
		for range 1000 {
			n, err := watcher.Write([]byte("Test"))

			require.NoError(t, err)
			require.Equal(t, n, 4)
		}

		watcher.Close()
	}()

	var actualBytesWritten int
	for p := range watcher.Notifier() {
		actualBytesWritten += p
	}

	require.Equal(t, actualBytesWritten, 4000)
	require.Equal(t, buf.Bytes(), bytes.Repeat([]byte("Test"), 1000))
}
