package session

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/maxcelant/pomo-cli/internal/manager"
	"github.com/maxcelant/pomo-cli/internal/screen"
	"github.com/maxcelant/pomo-cli/internal/state"
	"github.com/maxcelant/pomo-cli/internal/timer"
)

type TimerCallback func(currentTime int)

type Session struct {
	manager.StateManager
	timer     timer.Timer
	intervals int
	options   map[string]interface{}
  reader    *bufio.Reader
}

func New(sm manager.StateManager, t timer.Timer, i int, r *bufio.Reader) *Session {
	return &Session{sm, t, i, make(map[string]interface{}), r}
}

func (s *Session) Start(options map[string]interface{}) {
	s.options = options
	for {
		s.loop(state.ACTIVE)
		s.loop(state.REST)
		s.incrementInterval()
	}
}

func (s Session) handleInput(inputChan chan<- string) {
	for {
		input, _ := s.reader.ReadString('\n')
		inputChan <- strings.TrimSpace(input)
	}
}

func (s Session) loop(nextState state.ID) {
	inputChan := make(chan string)
	controlChan := make(chan struct{})
	go s.handleInput(inputChan)

	s.UpdateState(state.Get(nextState))
	s.timer.SetDuration(s.State.Duration)

	s.Time(nextState, inputChan, controlChan)
	s.awaitUserResponse()
}

func (s *Session) Time(nextState state.ID, inputChan chan string, controlChan chan struct{}) {
	out := make(chan int)

	go s.timer.Countdown(out, controlChan)

	for {
		select {
		case input := <-inputChan:
			if input == "" && s.State.Id == state.PAUSE {
				s.UpdateState(state.Get(nextState))
				controlChan <- struct{}{}
			} else if input == "" && s.State.Id != state.PAUSE {
				s.UpdateState(state.Get(state.PAUSE))
				screen.Clear()
				fmt.Printf("🍎 Timer paused!\nPress [Enter] to unpause: ")
				controlChan <- struct{}{}
			}
		case time, ok := <-out:
			if !ok {
				return
			}
			s.PromptText(time)
		}
	}
}

// todo: refactor this
func (s Session) PromptText(t int) {
	screen.Clear()
	min, sec := s.timer.FormatDuration(s.State.Duration - t)
	if s.options["minimal"] == true {
		fmt.Printf("🍎 State: %s %s %d:%d\n", s.State.Literal, s.State.Symbol, min, sec)
		fmt.Printf("   Press [Enter] to pause timer: ")
		return
	}
	fmt.Println("🍎 Time to focus")
	fmt.Printf("   State: %s %s\n", s.State.Literal, s.State.Symbol)
	fmt.Printf("   Interval: %d\n", s.intervals)
	fmt.Printf("   Time Remaining: %dm %ds\n", min, sec)
	fmt.Printf("   Press [Enter] to pause timer: ")
}

func (s *Session) awaitUserResponse() {
	screen.Clear()
	if s.State.Id == state.ACTIVE {
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
