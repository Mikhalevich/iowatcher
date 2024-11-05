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
		initialBytes          = bytes.Repeat([]byte("Test"), 1000)
		expectedBytesNotified = 4000
	)

	results, err := processRead(initialBytes)

	require.NoError(t, err)
	require.Equal(t, initialBytes, results.ActualBytesRead)
	require.Equal(t, expectedBytesNotified, results.ActualBytesNotified)
}

type readResults struct {
	ActualBytesRead     []byte
	ActualBytesNotified int
}

func processRead(initialBytes []byte) (readResults, error) {
	var (
		watcher = NewReadWatcher(bytes.NewReader(initialBytes))
		wg      sync.WaitGroup
		results readResults
	)

	wg.Add(1)

	go func() {
		defer wg.Done()

		for p := range watcher.Notifier() {
			results.ActualBytesNotified += p
		}
	}()

	actualBytesRead, err := io.ReadAll(watcher)
	results.ActualBytesRead = actualBytesRead

	wg.Wait()

	//nolint:wrapcheck
	return results, err
}

func BenchmarkReader100Bytes(b *testing.B) {
	var (
		initialBytes          = bytes.Repeat([]byte("100bt"), 20)
		expectedBytesNotified = 100
	)

	for i := 0; i < b.N; i++ {
		results, err := processRead(initialBytes)

		require.NoError(b, err)
		require.Equal(b, initialBytes, results.ActualBytesRead)
		require.Equal(b, expectedBytesNotified, results.ActualBytesNotified)
	}
}

func BenchmarkReader4KBytes(b *testing.B) {
	var (
		initialBytes          = bytes.Repeat([]byte("Test"), 1000)
		expectedBytesNotified = 4000
	)

	for i := 0; i < b.N; i++ {
		results, err := processRead(initialBytes)

		require.NoError(b, err)
		require.Equal(b, initialBytes, results.ActualBytesRead)
		require.Equal(b, expectedBytesNotified, results.ActualBytesNotified)
	}
}

func BenchmarkReader64KBytes(b *testing.B) {
	var (
		initialBytes          = bytes.Repeat([]byte("Test"), 16000)
		expectedBytesNotified = 64000
	)

	for i := 0; i < b.N; i++ {
		results, err := processRead(initialBytes)

		require.NoError(b, err)
		require.Equal(b, initialBytes, results.ActualBytesRead)
		require.Equal(b, expectedBytesNotified, results.ActualBytesNotified)
	}
}
