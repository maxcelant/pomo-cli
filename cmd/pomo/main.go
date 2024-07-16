package main

import (
	"fmt"
	"os"
	"time"
)

type State int

const (
  INIT State = iota
  ACTIVE
  BREAK
  DONE
)

var activeTime int = 25
var breakTime int = 5
var intervals int = 3 

func usage() {
	fmt.Println("Usage: pomo [command]")
}

func timer(timeAmount int) {
	for t := 0; t < timeAmount; t++ {
		select {
		case <-time.After(1 * time.Second):
			fmt.Printf("%d remaining\n", timeAmount-t)
		}
	}
}

func handleStartCommand(subcommands []string) {
  curState := INIT

	for _, s := range subcommands {
		fmt.Printf("%s\n", s)
	}
  
  for i := 0; i < intervals; i++ {
    curState = ACTIVE
    fmt.Printf("State is now active...%d\n", curState)
	  timer(activeTime)
    curState = BREAK
    fmt.Printf("State is now break...%d\n", curState)
	  timer(breakTime)
  }
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
