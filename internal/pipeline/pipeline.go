package pipeline

import (
	"bufio"
	"io"

	"github.com/logpipe/internal/filter"
	"github.com/logpipe/internal/transform"
)

// Pipeline reads lines from a source, applies filters and a formatter, and writes to a sink.
type Pipeline struct {
	filter    *filter.Filter
	formatter *transform.Formatter
	reader    io.Reader
	writer    io.Writer
}

// Config holds options for constructing a Pipeline.
type Config struct {
	Filter    *filter.Filter
	Formatter *transform.Formatter
	Reader    io.Reader
	Writer    io.Writer
}

// New creates a new Pipeline from the given Config.
func New(cfg Config) (*Pipeline, error) {
	if cfg.Reader == nil {
		return nil, fmt.Errorf("pipeline: reader must not be nil")
	}
	if cfg.Writer == nil {
		return nil, fmt.Errorf("pipeline: writer must not be nil")
	}
	return &Pipeline{
		filter:    cfg.Filter,
		formatter: cfg.Formatter,
		reader:    cfg.Reader,
		writer:    cfg.Writer,
	}, nil
}

// Run processes the input stream line by line until EOF or an error.
func (p *Pipeline) Run() error {
	scanner := bufio.NewScanner(p.reader)
	for scanner.Scan() {
		line := scanner.Text()
		if p.filter != nil && !p.filter.Keep(line) {
			continue
		}
		var out string
		if p.formatter != nil {
			var err error
			out, err = p.formatter.Format(line)
			if err != nil {
				return err
			}
		} else {
			out = line
		}
		if _, err := io.WriteString(p.writer, out+"\n"); err != nil {
			return err
		}
	}
	return scanner.Err()
}
