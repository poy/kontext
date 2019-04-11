package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/poy/kontext"
)

var (
	directory = flag.String("directory", "", "The directory to upload as context.")
	tag       = flag.String("tag", "", "The tag to which we upload the context.")
	rebase    = flag.Bool("rebase", false, "Rebase the source.")
)

func main() {
	flag.Parse()

	if *directory == "" {
		log.Fatalf("Missing required flag: --directory")
	}

	if *tag == "" {
		log.Fatalf("Missing required flag: --tag")
	}

	err := kontext.BuildImage(*directory, *tag, *rebase)
	if err == kontext.ErrNoChange {
		fmt.Fprintln(os.Stderr, "no change in source context (or empty)")
		return
	}

	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
