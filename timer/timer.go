package timer

import "time"

type TimerSubscriber interface {
	OnTimer()
}

type Timer interface {
	Subscribe(subscriber TimerSubscriber)
	Start()
}

func NewTimer(duration time.Duration) Timer {
	return &timedTimer{
		duration:    duration,
		subscribers: make([]TimerSubscriber, 0),
	}
}

type timedTimer struct {
	duration    time.Duration
	subscribers []TimerSubscriber
}

func (t *timedTimer) Subscribe(subscriber TimerSubscriber) {
	t.subscribers = append(t.subscribers, subscriber)
}

func (t *timedTimer) Start() {
	realTimer := time.NewTimer(t.duration)
	go func() {
		<-realTimer.C
		for _, subscriber := range t.subscribers {
			subscriber.OnTimer()
		}

		realTimer.Stop()
	}()
}
