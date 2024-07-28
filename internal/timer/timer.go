package timer

import (
  "time"
  "fmt"
  "bufio"
  "os"
  "github.com/maxcelant/pomo-cli/internal/screen"
)

type TimerCallback func(currentTime int)

type Timer struct {
	duration int
  pauseChan chan struct{}
}

func New() Timer {
	return Timer{0, make(chan struct{})}
}

func (t Timer) countdown(out chan<- int) {
	for i := 0; i < t.duration; i++ {
		select {
		case <-time.After(1 * time.Second):
			out <- i
    case <- t.pauseChan:
      screen.Clear()
      fmt.Println("🍎 Session Paused")
      <-t.pauseChan
		}
	}
	close(out)
}

func (t Timer) listenForPause(reader *bufio.Reader) {
  for {
    _, _ = reader.ReadString('\n')
    t.Swap()
  }
}

func (t Timer) Time(cb TimerCallback) {
	reader := bufio.NewReader(os.Stdin)
  out := make(chan int)
  
  go t.countdown(out)
  go t.listenForPause(reader)

	for time := range out {
		cb(time)
	}
}

func (t *Timer) Swap() {
  t.pauseChan<-struct{}{}
}

func (t *Timer) SetDuration(duration int) {
	t.duration = duration
}

func (t Timer) FormatDuration(duration int) (int, int) {
	minutes := duration / 60
	seconds := duration % 60
	return minutes, seconds
}
