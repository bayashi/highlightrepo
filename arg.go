package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	flag "github.com/spf13/pflag"
)

const (
	cmdName string = "highlightrepo"

	exitOK  int = 0
	exitErr int = 1
)

var (
	version     = ""
	installFrom = "Source"
)

type options struct {
	color   string
	noTilde bool
}

func (cli *runner) parseArgs() *options {
	o := &options{}

	var flagHelp bool
	var flagVersion bool
	flag.StringVarP(&o.color, "color", "c", "cyan", "Color name to highlight")
	flag.BoolVarP(&o.noTilde, "no-tilde", "t", false, "Not replace home dir path to tilda")
	flag.BoolVarP(&flagHelp, "help", "h", false, "Show help (This message) and exit")
	flag.BoolVarP(&flagVersion, "version", "v", false, "Show version and build command info and exit")

	flag.CommandLine.SortFlags = false
	flag.Parse()

	if flagHelp {
		cli.putHelp(fmt.Sprintf("Version %s", getVersion()))
	}

	if flagVersion {
		cli.putErr(versionDetails())
		os.Exit(exitOK)
	}

	return o
}

func versionDetails() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH
	compiler := runtime.Version()

	return fmt.Sprintf(
		"Version %s - %s.%s (compiled:%s, %s)",
		getVersion(),
		goos,
		goarch,
		compiler,
		installFrom,
	)
}

func getVersion() string {
	if version != "" {
		return version
	}
	i, ok := debug.ReadBuildInfo()
	if !ok {
		return "Unknown"
	}

	return i.Main.Version
}

func (cli *runner) putErr(message ...interface{}) {
	fmt.Fprintln(cli.err, message...)
}

func (cli *runner) putUsage() {
	cli.putErr(fmt.Sprintf("Usage: pwd | %s [OPTIONS]", cmdName))
}

func (cli *runner) putHelp(message string) {
	cli.putErr(message)
	cli.putUsage()
	cli.putErr("Options:")
	flag.PrintDefaults()
	os.Exit(exitOK)
}
