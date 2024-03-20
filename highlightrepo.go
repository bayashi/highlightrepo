package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/bayashi/colorpalette"
	"github.com/fatih/color"
)

const GITDIR = ".git"
const TILDE = "~"

var SEP = string(os.PathSeparator)

type highlightrepo struct {
	options          *options
	hlColor          *color.Color
	highlightedLines []string
}

func NewHighlightrepo(o *options) *highlightrepo {
	h := &highlightrepo{
		options: o,
	}

	if o.nonTTY {
		color.NoColor = false
	}

	h.hlColor = colorpalette.Get(o.color)

	return h
}

func (h *highlightrepo) scan(cli *runner, scanner *bufio.Scanner) error {
	for scanner.Scan() {
		line := scanner.Bytes()

		if err := h.highlightline(string(line)); err != nil {
			return fmt.Errorf("error during Highlight %w", err)
		}

		if err := h.show(cli.out); err != nil {
			return fmt.Errorf("error during Show %w", err)
		}
	}

	return nil
}

func (h *highlightrepo) highlightline(line string) error {
	line = strings.Trim(line, " ")
	absPath, err := filepath.Abs(line)
	if err != nil {
		return err
	}

	if _, err := os.Stat(absPath); err != nil {
		return fmt.Errorf("orig %s, abs %s : %w", line, absPath, err)
	}

	paths := strings.Split(absPath, SEP)

	curPath := ""
	highlightedPath := ""
	for i, p := range paths {
		curPath = curPath + p + SEP
		if hasGitDir(curPath) {
			p = h.hlColor.Sprint(p)
		}
		highlightedPath = highlightedPath + p
		if len(paths) != i+1 {
			highlightedPath = highlightedPath + SEP
		}
	}

	cur, err := user.Current()
	if err != nil {
		return err
	}
	homeDir := cur.HomeDir

	// Don't put git directory within home directory path.
	// If there would be, not replace home directory path to a tilde automatically even without `--no-tilde` option.
	if !h.options.noTilde && strings.HasPrefix(highlightedPath, homeDir) {
		highlightedPath = strings.Replace(highlightedPath, homeDir, TILDE, 1)
	}

	h.highlightedLines = append(h.highlightedLines, highlightedPath)

	return nil
}

func hasGitDir(dirPath string) bool {
	d, err := os.Stat(dirPath + GITDIR)
	return err == nil && d.IsDir()
}

func (h *highlightrepo) show(w io.Writer) error {
	writer := bufio.NewWriter(w)
	for i, l := range h.highlightedLines {
		if _, err := fmt.Fprint(writer, l); err != nil {
			return err
		}
		if i+1 != len(h.highlightedLines) {
			if _, err := fmt.Fprint(writer, "\n"); err != nil {
				return err
			}
		}
		if err := writer.Flush(); err != nil {
			return err
		}
	}

	return nil
}
