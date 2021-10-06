package main

import (
	"os"
	"fmt"
	"flag"
    "runtime"
)

var Version = "development"

func main() {
	version := flag.Bool("version", false, "Prints current app version")
	goos := flag.Bool("goos", false, "Get GOOS")
	goarch := flag.Bool("goarch", false, "Get GOARCH")
	flag.Parse()
	if *version {
		fmt.Println(Version)
		return
	}

	if *goos {
		_, _ = os.Stdout.WriteString(runtime.GOOS)
	}

	if *goarch {
		_, _ = os.Stdout.WriteString(runtime.GOARCH)
	}
}
