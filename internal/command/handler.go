
package command

import (
	"fmt"
	"os"

	"github.com/maxcelant/pomo-cli/internal/fileio"
	"github.com/maxcelant/pomo-cli/internal/session"
	"github.com/maxcelant/pomo-cli/internal/subcommand"
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

func NewHandler(commandName string, session *session.Session, subcommands []string) (Handler, error) {
	switch commandName {
	case "start", "session":
		return NewStartCommandHandler(session, subcommands), nil
	case "config":
		return NewConfigCommandHandler(subcommands), nil
	default:
		return nil, fmt.Errorf("invalid command: %s", commandName)
	}
}
