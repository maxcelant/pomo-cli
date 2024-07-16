package main

import (
	"fmt"
	"os"
	"time"
)

var activeTime int = 25
var breakTime int = 5
var intervals int = 0

func usage() {
	fmt.Println("Usage: pomo [command]")
}

func timer(timeAmount int) {
  for t := 0; t < timeAmount; t++ {
    select {
    case <- time.After(1 * time.Second):
      fmt.Printf("%d remaining\n", timeAmount - t)
    }
  }
}

func handleStartCommand(subcommands []string) {
	for _, s := range subcommands {
		fmt.Printf("%s\n", s)
	}

	timer(activeTime)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		usage()
		os.Exit(1)
	}

	for _, arg := range args[1:] {
		if arg == "start" {
			handleStartCommand(args[2:])
			break
		}
	}
}
