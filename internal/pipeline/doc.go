// Package pipeline wires together a filter and a formatter to process
// a streaming line-oriented log source.
//
// Basic usage:
//
//	p, err := pipeline.New(pipeline.Config{
//		Reader:    os.Stdin,
//		Writer:    os.Stdout,
//		Filter:    myFilter,
//		Formatter: myFormatter,
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	if err := p.Run(); err != nil {
//		log.Fatal(err)
//	}
package pipeline
