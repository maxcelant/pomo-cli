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

func stopSignal(stopChan chan<- struct{}) {
	reader := bufio.NewReader(os.Stdin)
  for {
    input, _ := reader.ReadString('\n')
    if input == "s\n" {
      stopChan <- struct{}{}
      return
    }
  }
}

func (s Session) loop(nextState state.ID) {
  stopChan := make(chan struct{})
  go stopSignal(stopChan)
	
  s.UpdateState(state.Get(nextState))
	s.timer.SetDuration(s.State.Duration)
	s.timer.Time(stopChan, func(t int) {
    if s.options["silent"] == true {
      return 
    }
    screen.Clear()
    fmt.Println("🍎 Time to focus")
    fmt.Printf("   State: %s %s\n", s.State.Literal, s.State.Symbol)
    fmt.Printf("   Interval: %d\n", s.intervals)
    m, s := s.timer.FormatDuration(s.State.Duration-t)
    fmt.Printf("   Time Remaining: %dm %ds\n", m, s)
    fmt.Printf("   Press [s] to stop timer: ")
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
