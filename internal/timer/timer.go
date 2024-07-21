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

func (t Timer) Time(stopChan <-chan struct{}, cb TimerCallback) {
	out := make(chan int)

	go t.countdown(out)

	for {
		select {
		case <-stopChan:
			return
		case time, ok := <-out:
			if !ok {
				return
			}
			cb(time)
		}
	}
}

func (t *Timer) SetDuration(duration int) {
	t.duration = duration
}

func (t Timer) FormatDuration(duration int) (int, int) {
  minutes := duration / 60
  seconds := duration % 60
  return minutes, seconds
}
