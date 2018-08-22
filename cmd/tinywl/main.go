package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/alexbakker/go-wlroots/wlroots"
)

var (
	command = flag.String("s", "", "startup command")
)

func fatal(msg string, err error) {
	fmt.Printf("error %s: %s\n", msg, err)
	os.Exit(1)
}

func main() {
	flag.Parse()

	// set up logging
	wlroots.OnLog(wlroots.LogImportanceDebug, nil)

	// start the server
	server, err := NewServer()
	if err != nil {
		fatal("initializing server", err)
	}
	if err = server.Start(); err != nil {
		fatal("starting server", err)
	}

	// run the startup command if given
	if *command != "" {
		cmd := exec.Command("/bin/sh", "-c", *command)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err = cmd.Start(); err != nil {
			fatal("running startup command", err)
		}
	}

	// start the wayland event loop
	if err = server.Run(); err != nil {
		fatal("running server", err)
	}
}
