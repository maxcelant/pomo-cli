
package session

import "fmt"
import "os"
import "bufio"
import "strings"
import "github.com/maxcelant/pomo-cli/cmd/manager"
import "github.com/maxcelant/pomo-cli/cmd/timer"
import "github.com/maxcelant/pomo-cli/cmd/state"
import "github.com/maxcelant/pomo-cli/cmd/screen"

type Session struct {
  sm manager.StateManager
  timer timer.Timer
  intervals int
}

func New(sm manager.StateManager, t timer.Timer, i int) *Session {
  return &Session{sm, t, i}
}

func (s Session) Loop(nextState state.ID) {
  s.sm.UpdateState(state.Get(nextState))
  s.timer.SetDuration(s.sm.State.Duration)
  s.timer.Time(func (t int) {
    screen.Clear()
    fmt.Println("🍎 Time to focus")
    fmt.Printf("   State: %s %s\n", s.sm.State.Literal, s.sm.State.Symbol)
    fmt.Printf("   Interval: %d\n", s.intervals)
    fmt.Printf("   Time Remaining: %ds\n", s.sm.State.Duration-t)
  })
  s.awaitUserRes()
}

func (s *Session) awaitUserRes() {
  screen.Clear()
  if s.sm.State.Id == state.ACTIVE {
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

func (s *Session) IncrementInterval() {
  s.intervals++
}
