package main

import (
	"flag"
	"log"
	"os/exec"
)


var (
	path = flag.String("path", "", "The path to plugin file.")
	outpath = flag.String("destination", "", "Output file.")
)

func main() {
	flag.Parse()

	if *path == "" {
		log.Fatal("-path should be specified.")
		return
	}
	if *outpath == "" {
		log.Fatal("-destination should be specified.")
	}

	exec.Command("go build -buildmode=plugin -trimpath -o "+ *path).Run()
}
