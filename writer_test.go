package iowatcher

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWrite4KBytes(t *testing.T) {
	t.Parallel()

	var (
		buf                     = bytes.Buffer{}
		expectedBytesWrittenLen = 4000
		expectedBytesWritten    = bytes.Repeat([]byte("Test"), 1000)
		actualBytesNotified     int
		watcher                 = NewWriteWatcher(&buf, func(bytesProcessed int) {
			actualBytesNotified += bytesProcessed
		})
	)

	for range 1000 {
		n, err := watcher.Write([]byte("Test"))

		require.NoError(t, err)
		require.Equal(t, 4, n)
	}

	require.Equal(t, expectedBytesWrittenLen, actualBytesNotified)
	require.Equal(t, expectedBytesWritten, buf.Bytes())
}

type writeResults struct {
	ActualBytesWrite    []byte
	ActualBytesNotified int
}

func processWrite(bytesToWrite []byte, writesCount int) writeResults {
	var (
		buf     = bytes.Buffer{}
		results writeResults
		watcher = NewWriteWatcher(&buf, func(bytesProcessed int) {
			results.ActualBytesNotified += bytesProcessed
		})
	)

	for range writesCount {
		//nolint:errcheck
		watcher.Write(bytesToWrite)
	}

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
