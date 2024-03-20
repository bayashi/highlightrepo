# highlightrepo

<a href="https://github.com/bayashi/highlightrepo/actions" title="highlightrepo CI"><img src="https://github.com/bayashi/highlightrepo/workflows/main/badge.svg" alt="highlightrepo CI"></a>
<a href="https://goreportcard.com/report/github.com/bayashi/highlightrepo" title="highlightrepo report card" target="_blank"><img src="https://goreportcard.com/badge/github.com/bayashi/highlightrepo" alt="highlightrepo report card"></a>
<a href="https://pkg.go.dev/github.com/bayashi/highlightrepo" title="Go highlightrepo package reference" target="_blank"><img src="https://pkg.go.dev/badge/github.com/bayashi/highlightrepo.svg" alt="Go Reference: highlightrepo"></a>

`highlightrepo` provides a filter to highlight repository directory in a path string.

## Usage

```sh
$ pwd | highlightrepo
```

Used cyan to highlight. If you want to set another color, then add `--color` option.

```sh
$ pwd | highlightrepo --color="red"
```

Here's [color palette](https://github.com/bayashi/colorpalette/blob/main/colorpalette.go).

### For PS1

Use `-y` option in `PS1` to bypass the check for non-tty output streams

```sh
PS1="\u \$(pwd | highlightrepo -y)\n\$ "
```
![PS1 example](https://github.com/bayashi/highlightrepo/assets/42190/d7f2ad43-86bd-40cf-a1be-99611d326450)

## Installation

### Mac

```sh
brew tap bayashi/tap
brew install bayashi/tap/highlightrepo
```

### Binary install

Download binary from here: https://github.com/bayashi/highlightrepo/releases

### Go manual

If you have golang environment:

```cmd
go install github.com/bayashi/highlightrepo@latest
```

## License

MIT License

## Author

Dai Okabayashi: https://github.com/bayashi
