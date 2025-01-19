package benchmarks

import (
	"testing"

	"github.com/srg77global/concurrentDataProcessing/internal/processor"
)

func BenchmarkProcessFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = processor.ProcessFile("../testdata/largefile.txt", 5)
	}
}
