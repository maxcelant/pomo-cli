
package timer

import "time"
import "fmt"
import "github.com/maxcelant/pomo-cli/cmd/manager"
import "github.com/maxcelant/pomo-cli/cmd/screen"

type Timer struct {
  duration int
}

func New() *Timer {
  return &Timer{0}
}

func (t *Timer) countdown(out chan<- int) {
  for i := 0; i < t.duration; i++ {
    select {
    case <-time.After(1 * time.Second):
      out<-i
    }
  }
  close(out)
}

func (t *Timer) Time(sm *manager.StateManager) {
  out := make(chan int)

  go t.countdown(out)

  for time := range out {
    screen.Clear()
    fmt.Println("🍎 Time to focus")
    fmt.Printf("   State: %s %s\n", sm.State.Literal, sm.State.Symbol)
    fmt.Printf("   Interval: %d\n", sm.Intervals)
    fmt.Printf("   Time Remaining: %ds\n", sm.State.Duration-time)
  }
}

func (t *Timer) SetDuration(duration int) {
  t.duration = duration
}
