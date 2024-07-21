package timer

import "time"


type Timer struct {
	duration int
}

func New() Timer {
	return Timer{0}
}

func (t Timer) Countdown(out chan<- int) {
	for i := 0; i < t.duration; i++ {
		select {
		case <-time.After(1 * time.Second):
			out <- i
		}
	}
	close(out)
}


func (t *Timer) SetDuration(duration int) {
	t.duration = duration
}

func (t Timer) FormatDuration(duration int) (int, int) {
  minutes := duration / 60
  seconds := duration % 60
  return minutes, seconds
}
