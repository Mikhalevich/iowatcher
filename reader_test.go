package iowatcher

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRead4KBytes(t *testing.T) {
	var (
		buf     = bytes.Repeat([]byte("Test"), 1000)
		watcher = NewReadWatcher(bytes.NewReader(buf))
	)

	go func() {
		readBuf, err := io.ReadAll(watcher)
		require.NoError(t, err)
		require.Equal(t, buf, readBuf)
	}()

	var actualBytesRead int
	for p := range watcher.Notifier() {
		actualBytesRead += p
	}

	require.Equal(t, actualBytesRead, 4000)
}
