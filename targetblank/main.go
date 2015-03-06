package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	help = flag.Bool("h", false, "")
)

var linkPattern = regexp.MustCompile(`\[([^\]]+)\]\(([^\)]+)\)`)

func usage() {
	fmt.Println(`[text](url) -> <a href="url" target="_blank">text</a>
    cat input.md | targetblank > output.md`)
}

func main() {
	flag.Parse()

	if *help {
		usage()
		os.Exit(1)
	}

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	src, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	dst := linkPattern.ReplaceAll(src, []byte(`<a href="$2" target="_blank">$1</a>`))
	matches := linkPattern.FindAllSubmatch(src, -1)
	for _, match := range matches {
		fmt.Fprintf(os.Stderr, "Replacing %s\n", string(match[0]))
	}

	fmt.Print(string(dst))
	return nil
}
