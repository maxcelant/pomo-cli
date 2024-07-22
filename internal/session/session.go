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
	stateManager manager.StateManager
	timer        timer.Timer
	intervals    int
	options      map[string]interface{}
	reader       *bufio.Reader
}

func New(stateManager manager.StateManager, timer timer.Timer, intervals int, reader *bufio.Reader) *Session {
	return &Session{
		stateManager: stateManager,
		timer:        timer,
		intervals:    intervals,
		options:      make(map[string]interface{}),
		reader:       reader,
	}
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
	s.stateManager.UpdateState(state.Get(nextState))
	s.timer.SetDuration(s.stateManager.Duration)
	s.timer.Time(func(t int) {
		if s.options["silent"] == true {
			return
		}
		screen.Clear()
		fmt.Println("🍎 Time to focus")
		fmt.Printf("   State: %s %s\n", s.stateManager.Literal, s.stateManager.Symbol)
		fmt.Printf("   Interval: %d\n", s.intervals)
		m, s := s.timer.FormatDuration(s.stateManager.Duration - t)
		fmt.Printf("   Time Remaining: %dm %ds\n", m, s)
	})
	s.awaitUserResponse()
}

func (s *Session) awaitUserResponse() {
	screen.Clear()
	if s.stateManager.Id == state.ACTIVE {
		fmt.Printf("🍎 Active session done! Ready to start break?")
	} else {
		fmt.Printf("🍎 Break session done! Ready to start studying?")
	}
	fmt.Printf("   [Enter] to continue, or [Q]uit: ")
	input, _ := s.reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "q" || input == "Q" {
		fmt.Printf("  Exiting gracefully...")
		os.Exit(0)
	}
}

func (s *Session) incrementInterval() {
	s.intervals++
}
