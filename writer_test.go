package iowatcher

import (
	"bytes"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrite4KBytes(t *testing.T) {
	t.Parallel()

	var (
		buf                     = bytes.Buffer{}
		watcher                 = NewWriteWatcher(&buf)
		expectedBytesWrittenLen = 4000
		expectedBytesWritten    = bytes.Repeat([]byte("Test"), 1000)
		actualBytesWritten      int
		wg                      sync.WaitGroup
	)

	wg.Add(1)

	go func() {
		defer wg.Done()

		for p := range watcher.Notifier() {
			actualBytesWritten += p
		}
	}()

	for range 1000 {
		n, err := watcher.Write([]byte("Test"))

		require.NoError(t, err)
		require.Equal(t, 4, n)
	}

	watcher.Close()
	wg.Wait()

	require.Equal(t, expectedBytesWrittenLen, actualBytesWritten)
	require.Equal(t, expectedBytesWritten, buf.Bytes())
}
