package pipeline

import (
	"bufio"
	"io"

	"github.com/user/logpipe/internal/transform"
)

// NewDedupeStage returns a pipeline stage that removes consecutive duplicate
// lines from the stream. Non-consecutive duplicates are preserved.
func NewDedupeStage(r io.Reader, w io.Writer) error {
	if r == nil {
		return ErrNilReader
	}
	if w == nil {
		return ErrNilWriter
	}

	deduper := transform.NewDeduplicator()
	scanner := bufio.NewScanner(r)
	bw := bufio.NewWriter(w)

	for scanner.Scan() {
		line := scanner.Text()
		if out, ok := deduper.Format(line); ok {
			if _, err := bw.WriteString(out + "\n"); err != nil {
				return err
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return bw.Flush()
}
