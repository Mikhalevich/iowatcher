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

type writeResults struct {
	ActualBytesWrite    []byte
	ActualBytesNotified int
}

func processWrite(bytesToWrite []byte, writesCount int) writeResults {
	var (
		buf     = bytes.Buffer{}
		watcher = NewWriteWatcher(&buf)
		wg      sync.WaitGroup
		results writeResults
	)

	wg.Add(1)

	go func() {
		defer wg.Done()

		for p := range watcher.Notifier() {
			results.ActualBytesNotified += p
		}
	}()

	for range writesCount {
		//nolint:errcheck
		watcher.Write(bytesToWrite)
	}

	watcher.Close()
	wg.Wait()

	results.ActualBytesWrite = buf.Bytes()

	return results
}

func BenchmarkWrite100Bytes(b *testing.B) {
	var (
		bytesToWrite          = []byte("100bt")
		writesCount           = 20
		expectedBytesNotified = 100
	)

	for i := 0; i < b.N; i++ {
		results := processWrite(bytesToWrite, writesCount)

		b.StopTimer()

		require.Equal(b, bytes.Repeat(bytesToWrite, writesCount), results.ActualBytesWrite)
		require.Equal(b, expectedBytesNotified, results.ActualBytesNotified)

		b.StartTimer()
	}
}

func BenchmarkWriter4KBytes(b *testing.B) {
	var (
		bytesToWrite          = []byte("Test")
		writesCount           = 1000
		expectedBytesNotified = 4000
	)

	for i := 0; i < b.N; i++ {
		results := processWrite(bytesToWrite, writesCount)

		b.StopTimer()

		require.Equal(b, bytes.Repeat(bytesToWrite, writesCount), results.ActualBytesWrite)
		require.Equal(b, expectedBytesNotified, results.ActualBytesNotified)

		b.StartTimer()
	}
}

func BenchmarkWriter64KBytes(b *testing.B) {
	var (
		bytesToWrite          = []byte("Test")
		writesCount           = 16000
		expectedBytesNotified = 64000
	)

	for i := 0; i < b.N; i++ {
		results := processWrite(bytesToWrite, writesCount)

		b.StopTimer()

		require.Equal(b, bytes.Repeat(bytesToWrite, writesCount), results.ActualBytesWrite)
		require.Equal(b, expectedBytesNotified, results.ActualBytesNotified)

		b.StartTimer()
	}
}
