package iowatcher

import (
	"bytes"
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRead4KBytes(t *testing.T) {
	t.Parallel()

	var (
		buf                  = bytes.Repeat([]byte("Test"), 1000)
		watcher              = NewReadWatcher(bytes.NewReader(buf))
		expectedBytesReadLen = 4000
		actualBytesRead      int
		wg                   sync.WaitGroup
	)

	wg.Add(1)

	go func() {
		defer wg.Done()

		for p := range watcher.Notifier() {
			actualBytesRead += p
		}
	}()

	actualBuf, err := io.ReadAll(watcher)

	require.NoError(t, err)
	require.Equal(t, buf, actualBuf)

	wg.Wait()

	require.Equal(t, expectedBytesReadLen, actualBytesRead)
}
