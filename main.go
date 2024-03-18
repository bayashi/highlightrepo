package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"syscall"

	"golang.org/x/term"
)

type runner struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

func main() {
	cli := &runner{
		in:  os.Stdin,
		out: os.Stdout,
		err: os.Stderr,
	}
	cli.run()

	os.Exit(exitOK)
}

func (cli *runner) run() {
	o := cli.parseArgs()

	if term.IsTerminal(int(syscall.Stdin)) {
		os.Exit(exitOK)
	}

	err := cli.runner(o)
	if err != nil {
		cli.putErr(fmt.Sprintf("%s", err))
		os.Exit(exitErr)
	}
}

func (cli *runner) runner(o *options) error {
	h := NewHighlightrepo(o)
	if err := h.scan(cli, bufio.NewScanner(cli.in)); err != nil {
		return fmt.Errorf("error during scan: %w", err)
	}

	return nil
}
