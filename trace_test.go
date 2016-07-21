package trace_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tj/go-trace"
)

func TestTrace_Start_withoutArgs(t *testing.T) {
	var buf bytes.Buffer
	probe := trace.Trace{ID: "foo", Writer: &buf}
	probe.Start("db.query", nil)
	assert.Equal(t, `>:foo:db.query::`, buf.String())
}

func TestTrace_Stop_withoutArgs(t *testing.T) {
	var buf bytes.Buffer
	probe := trace.Trace{ID: "foo", Writer: &buf}
	probe.Stop("db.query", nil)
	assert.Equal(t, `<:foo:db.query::`, buf.String())
}

func TestTrace_Start_withArgs(t *testing.T) {
	var buf bytes.Buffer
	probe := trace.Trace{ID: "foo", Writer: &buf}
	probe.Start("db.query", trace.Args{"user": "tobi", "species": "ferret"})
	assert.Equal(t, `>:foo:db.query:user=tobi,species=ferret:`, buf.String())
}

func TestTrace_Stop_withArgs(t *testing.T) {
	var buf bytes.Buffer
	probe := trace.Trace{ID: "foo", Writer: &buf}
	probe.Stop("db.query", trace.Args{"user": "tobi", "species": "ferret"})
	assert.Equal(t, `<:foo:db.query:user=tobi,species=ferret:`, buf.String())
}

func TestTrace_Stop_withNumericValues(t *testing.T) {
	var buf bytes.Buffer
	probe := trace.Trace{ID: "foo", Writer: &buf}
	probe.Stop("db.query", trace.Args{"user": "tobi", "age": 3})
	assert.Equal(t, `<:foo:db.query:user=tobi,age=3:`, buf.String())
}

func TestTrace_Stop_escaped(t *testing.T) {
	var buf bytes.Buffer
	probe := trace.Trace{ID: "foo", Writer: &buf}
	probe.Stop("db.query", trace.Args{"user": "tobi.ferret"})
	assert.Equal(t, `<:foo:db.query:user=tobi\.ferret:`, buf.String())
}

func TestEscape(t *testing.T) {
	input := "foo.bar.baz,test="
	output := trace.Escape(input)
	assert.Equal(t, `foo\.bar\.baz\,test\=`, output)
}
