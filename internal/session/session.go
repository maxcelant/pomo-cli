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
	pauseChan    chan struct{}
}

func New(stateManager manager.StateManager, timer timer.Timer, intervals int, reader *bufio.Reader) *Session {
	return &Session{
		stateManager: stateManager,
		timer:        timer,
		intervals:    intervals,
		options:      make(map[string]interface{}),
		reader:       reader,
		pauseChan:    make(chan struct{}),
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

	go s.timer.Time(s.pauseChan, func(t int) {
		if s.timer.Duration <= 0 {
			s.stateManager.UpdateState(state.Get(state.WAITING))
		}
		if s.options["silent"] == true || t%5 != 0 {
			return
		}
		screen.Clear()
		fmt.Println("🍎 Time to focus")
		fmt.Printf("   State: %s %s\n", s.stateManager.Literal, s.stateManager.Symbol)
		if s.options["intervals"] != -1 {
			fmt.Printf("   Interval: %d/%d\n", s.intervals, s.options["intervals"])
		} else {
			fmt.Printf("   Interval: %d\n", s.intervals)
		}
		m, s := s.timer.FormatDuration(s.stateManager.Duration - t)
		fmt.Printf("   Time Remaining: %dm %ds\n", m, s)
		fmt.Print("   Press [Enter] to pause or [q] to quit")
	})

	s.handlePrompt()
}

func (s *Session) handlePrompt() {
	for {
		input, _ := s.reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "q" || input == "Q" {
			fmt.Printf("   Exiting gracefully...")
			os.Exit(0)
		}
		if s.stateManager.Id == state.WAITING {
			return
		}
		s.swap()
	}
}

func (s *Session) swap() {
	s.pauseChan <- struct{}{}
}

func (s *Session) incrementInterval() {
	s.intervals++
}
