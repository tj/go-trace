// Package trace implements the Sysdig tracing /dev/null protocol.
package trace

import (
	"fmt"
	"io"
	"strings"
)

// Args for a tracer.
type Args map[string]interface{}

// Trace probe.
type Trace struct {
	ID     string    // ID of the trace
	Writer io.Writer // Writer which should be /dev/null in production
}

// Start tracer.
func (t *Trace) Start(tags string, args Args) error {
	_, err := t.Writer.Write([]byte(formatTrace(">", t.ID, tags, args)))
	return err
}

// Stop tracer.
func (t *Trace) Stop(tags string, args Args) error {
	_, err := t.Writer.Write([]byte(formatTrace("<", t.ID, tags, args)))
	return err
}

// format trace.
func formatTrace(dir, id, tags string, args Args) string {
	return fmt.Sprintf("%s:%s:%s:%s:", dir, id, tags, formatArgs(args))
}

// format arguments.
func formatArgs(args Args) string {
	var pairs []string

	for k, v := range args {
		pairs = append(pairs, formatArg(k, v))
	}

	return strings.Join(pairs, ",")
}

// format argument.
func formatArg(k string, v interface{}) string {
	return fmt.Sprintf("%s=%s", Escape(k), Escape(fmt.Sprintf("%v", v)))
}

// Escape string.
func Escape(s string) string {
	s = strings.Replace(s, ".", "\\.", -1)
	s = strings.Replace(s, ",", "\\,", -1)
	s = strings.Replace(s, "=", "\\=", -1)
	return s
}
