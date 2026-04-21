package transform

// Deduplicator filters out consecutive duplicate lines from a log stream.
type Deduplicator struct {
	last string
	seenAny bool
}

// NewDeduplicator creates a new Deduplicator that suppresses consecutive
// identical lines. Each call to Format returns the line unchanged if it
// differs from the previous line, or an empty string to signal suppression.
func NewDeduplicator() *Deduplicator {
	return &Deduplicator{}
}

// Format returns the line if it is different from the last seen line.
// If the line is a duplicate of the previous one, it returns ("", false)
// so callers know to skip it.
func (d *Deduplicator) Format(line string) (string, bool) {
	if d.seenAny && line == d.last {
		return "", false
	}
	d.last = line
	d.seenAny = true
	return line, true
}

// Reset clears the deduplication state, treating the next line as fresh.
func (d *Deduplicator) Reset() {
	d.last = ""
	d.seenAny = false
}
