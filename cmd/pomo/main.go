package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/maxcelant/pomo-cli/internal/manager"
	"github.com/maxcelant/pomo-cli/internal/screen"
	"github.com/maxcelant/pomo-cli/internal/session"
	"github.com/maxcelant/pomo-cli/internal/state"
	"github.com/maxcelant/pomo-cli/internal/timer"
)

func handleStartCommand(session *session.Session, subcommands []string) {
	for _, s := range subcommands {
		fmt.Printf("%s\n", s)
	}

	for {
		session.Loop(state.ACTIVE)
		session.Loop(state.REST)
		session.IncrementInterval()
	}
}

type Flag struct {
	datatype string
	name     string
}

func subcommandHandler(subcommands []string, out map[string]interface{}) (map[string]interface{}, error) {
	flags := map[string]Flag{
		"-a":       {datatype: "int", name: "active"},
		"--active": {datatype: "int", name: "active"},
		"-r":       {datatype: "int", name: "rest"},
		"--rest":   {datatype: "int", name: "rest"},
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error: input was not given for one of your flags")
			os.Exit(1)
		}
	}()

	for i := 0; i < len(subcommands); i++ {
		flag := subcommands[i]

		f, found := flags[flag]
		if !found {
			return nil, fmt.Errorf("flag '%s' is not a viable flag", flag)
		}

		if i+1 >= len(subcommands) {
			return nil, fmt.Errorf("flag '%s' expects a value but none was provided", flag)
		}

		if f.datatype != "int" {
			return nil, fmt.Errorf("datatype '%s' not implemented yet", f.datatype)
		}

		nextValue := subcommands[i+1]
		duration, err := strconv.Atoi(nextValue)
		if err != nil {
			return nil, fmt.Errorf("value for flag '%s' is not a valid integer: %s", flag, nextValue)
		}

		out[f.name] = duration
		i++
	}

	return out, nil
}

func handleConfigCommand(session *session.Session, subcommands []string) {
	out, err := subcommandHandler(subcommands, map[string]interface{}{"active": 0, "rest": 0})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Println("Subcommand values:", out)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		screen.Usage()
		os.Exit(1)
	}

	sm := manager.New(state.Get(state.INIT))
	timer := timer.New()
	session := session.New(sm, timer, 0)

	switch args[1] {
	case "start":
		handleStartCommand(session, args[2:])
	case "config":
		handleConfigCommand(session, args[2:])
	default:
		screen.Usage()
	}
}
