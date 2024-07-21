package session

import "fmt"
import "os"
import "bufio"
import "strings"
import "github.com/maxcelant/pomo-cli/internal/manager"
import "github.com/maxcelant/pomo-cli/internal/timer"
import "github.com/maxcelant/pomo-cli/internal/state"
import "github.com/maxcelant/pomo-cli/internal/screen"

type TimerCallback func(currentTime int)

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

func (s Session) handleInput(inputChan chan<- string) {
  reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		inputChan <- strings.TrimSpace(input)
	}
}

func (s Session) loop(nextState state.ID) {
  inputChan := make(chan string)
  go s.handleInput(inputChan)

  s.UpdateState(state.Get(nextState))
	s.timer.SetDuration(s.State.Duration)

	s.Time(nextState, inputChan, func(t int) {
    s.PromptText(t)
	})
	s.awaitUserResponse()
}

func (s *Session) Time(nextState state.ID, inputChan chan string, cb TimerCallback) {
	out := make(chan int)

	go s.timer.Countdown(out)

	for {
		select {
    case input := <-inputChan:
      if input == "" && s.State.Id == state.PAUSE {
        s.UpdateState(state.Get(nextState))
      } else if input == "" && s.State.Id != state.PAUSE {
        s.UpdateState(state.Get(state.PAUSE))
      }
		case time, ok := <-out:
			if !ok {
				return
			}
			cb(time)
		}
	}
}

// todo: refactor this
func (s Session) PromptText(t int) {
  screen.Clear()
  if s.options["silent"] == true {
    min, _ := s.timer.FormatDuration(s.State.Duration)
    fmt.Printf("🍎 State: %s %s %dm\n", s.State.Literal, s.State.Symbol, min)
    fmt.Printf("   Press [Enter] to pause timer: ")
    return 
  }
  if s.State.Id == state.PAUSE {
    fmt.Printf("🍎 Timer paused!\nPress [Enter] to unpause: ")
    return
  }
  fmt.Println("🍎 Time to focus")
  fmt.Printf("   State: %s %s\n", s.State.Literal, s.State.Symbol)
  fmt.Printf("   Interval: %d\n", s.intervals)
  min, sec := s.timer.FormatDuration(s.State.Duration-t)
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
