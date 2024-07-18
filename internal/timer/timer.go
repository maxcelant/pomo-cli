package timer

import "time"

type TimerCallback func(currentTime int)

type Timer struct {
	duration int
}

func New() Timer {
	return Timer{0}
}

func (t Timer) countdown(out chan<- int) {
	for i := 0; i < t.duration; i++ {
		select {
		case <-time.After(1 * time.Second):
			out <- i
		}
	}
	close(out)
}

func (t Timer) Time(cb TimerCallback) {
	out := make(chan int)

	go t.countdown(out)

	for time := range out {
		cb(time)
	}
}

func (t *Timer) SetDuration(duration int) {
	t.duration = duration
}
