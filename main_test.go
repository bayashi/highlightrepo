package main

import (
	"bytes"
	"strings"
	"testing"

	w "github.com/bayashi/witness"
	"github.com/fatih/color"
)

func TestRunner_OK(t *testing.T) {
	color.NoColor = false

	// `foo/` doesn't have any dirs or files in it.
	// `bar/` have `.git` file. But it's just a file. So NOT highlighted.
	// `baz/` have `.git` directory. It's expected to be highlighted.
	for tname, tt := range map[string]struct {
		in     string
		expect string
	}{
		"foo":                      {in: "testdata/foo", expect: "testdata/foo"},
		"bar":                      {in: "testdata/bar", expect: "testdata/bar"},
		"baz":                      {in: "testdata/baz", expect: "testdata/\x1b[96mbaz\x1b[0m"},
		"multiple lines not match": {in: "testdata/foo\ntestdata/bar", expect: "testdata/bar"},
		"multiple lines match":     {in: "testdata/bar\ntestdata/baz", expect: "testdata/\x1b[96mbaz\x1b[0m"},
	} {
		t.Run(tname, func(t *testing.T) {
			in := bytes.NewBufferString(tt.in)
			var out bytes.Buffer
			cli := &runner{
				in:  in,
				out: &out,
			}

			err := cli.runner(&options{color: "cyan"})

			if err != nil {
				w.Fail(t, "unexpected error", err)
			}

			g := out.String()
			if !strings.HasSuffix(g, tt.expect) {
				w.Fail(t, "Not match suffix", g, tt.expect)
			}
		})
	}
}
