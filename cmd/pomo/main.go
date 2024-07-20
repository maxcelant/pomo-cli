package main

import (
	"fmt"
	"os"

	"github.com/maxcelant/pomo-cli/internal/fileio"
	"github.com/maxcelant/pomo-cli/internal/manager"
	"github.com/maxcelant/pomo-cli/internal/screen"
	"github.com/maxcelant/pomo-cli/internal/session"
	"github.com/maxcelant/pomo-cli/internal/state"
	"github.com/maxcelant/pomo-cli/internal/subcommand"
	"github.com/maxcelant/pomo-cli/internal/timer"
)

type Handler interface {
  Handle() 
}

type StartCommandHandler struct {
  *session.Session
  subcommands []string
}

type ConfigCommandHandler struct {
  subcommands []string
}

func NewStartCommandHandler(session *session.Session, subcommands []string) *StartCommandHandler {
  return &StartCommandHandler{
    session,
    subcommands,
  }
}

func NewConfigCommandHandler(subcommands []string) *ConfigCommandHandler {
  return &ConfigCommandHandler{
    subcommands,
  }
}

func (s *StartCommandHandler) Handle() {
	options, err := subcommand.Handler(s.subcommands, map[string]interface{}{"silent": false})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
  s.Start(options)
}

func (c *ConfigCommandHandler) Handle() {
  options, err := subcommand.Handler(c.subcommands, map[string]interface{}{"active": 0, "rest": 0, "link":""})
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
  fileio.WriteToLocalYaml(options)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		screen.Usage()
		os.Exit(1)
	}

	sm := manager.New(state.Get(state.INIT))
	timer := timer.New()
	session := session.New(sm, timer, 1)

  handlers := map[string]Handler{
    "start": NewStartCommandHandler(session, args[2:]),
    "session": NewStartCommandHandler(session, args[2:]),
    "config": NewConfigCommandHandler(args[2:]),
  }

  handler, ok := handlers[args[1]]
  if !ok {
    fmt.Println("Error: invalid command")
    screen.Usage()
    os.Exit(1)
  }

  handler.Handle()
}
