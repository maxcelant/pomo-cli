package timer

import (
	"fmt"
	"github.com/maxcelant/pomo-cli/internal/screen"
	"time"
)

type TimerCallback func(currentTime int)

type Timer struct {
	Duration int
}

func New() Timer {
	return Timer{
		Duration: 0,
	}
}

func (t Timer) countdown(pauseChan chan struct{}, out chan<- int) {
	for i := 0; i < t.Duration; i++ {
		select {
		case <-time.After(1 * time.Second):
			out <- i
		case <-pauseChan:
			screen.Clear()
			fmt.Println("🍎 Session Paused")
			<-pauseChan
		}
	}
	close(out)
}

func (t Timer) Time(pauseChan chan struct{}, cb TimerCallback) {
	out := make(chan int)

	go t.countdown(pauseChan, out)

	for time := range out {
		cb(time)
	}
}

func (t *Timer) SetDuration(duration int) {
	t.Duration = duration
}

func (t Timer) FormatDuration(duration int) (int, int) {
	minutes := duration / 60
	seconds := duration % 60
	return minutes, seconds
}
