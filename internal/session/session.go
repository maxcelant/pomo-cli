package session

import "fmt"
import "os"
import "bufio"
import "strings"
import "github.com/maxcelant/pomo-cli/internal/manager"
import "github.com/maxcelant/pomo-cli/internal/timer"
import "github.com/maxcelant/pomo-cli/internal/state"
import "github.com/maxcelant/pomo-cli/internal/screen"

type Session struct {
	manager.StateManager
	timer     timer.Timer
	intervals int
  options   map[string]interface{}
}

func New(sm manager.StateManager, t timer.Timer, i int) *Session {
	return &Session{sm, t, i, make(map[string]interface{})}
}

func (s *Session) Start(options map[string]interface{}) {
  s.options = options
	for {
		s.loop(state.ACTIVE)
		s.loop(state.REST)
		s.incrementInterval()
	}
}

func (s Session) loop(nextState state.ID) {
	s.UpdateState(state.Get(nextState))
	s.timer.SetDuration(s.State.Duration)
	s.timer.Time(func(t int) {
    if s.options["silent"] == true {
      return 
    }
    screen.Clear()
    fmt.Println("🍎 Time to focus")
    fmt.Printf("   State: %s %s\n", s.State.Literal, s.State.Symbol)
    fmt.Printf("   Interval: %d\n", s.intervals)
    fmt.Printf("   Time Remaining: %ds\n", s.State.Duration-t)
	})
	s.awaitUserResponse()
}

func (s *Session) awaitUserResponse() {
	screen.Clear()
	if s.State.Id == state.ACTIVE {
		fmt.Printf("🍎 Active session done! Ready to start break?")
	} else {
		fmt.Printf("🍎 Break session done! Ready to start studying?")
	}
	fmt.Printf("   [Enter] to continue, or [Q]uit: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "q" || input == "Q" {
		fmt.Printf("  Exiting gracefully...")
		os.Exit(0)
	}
}

func (s *Session) incrementInterval() {
	s.intervals++
}
