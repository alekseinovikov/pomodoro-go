package timer

import "time"

type Timer interface {
	Periodic() Timer
	Subscribe(subscriber func()) Timer
	Start()
	Stop()
}

func NewTimer(duration time.Duration) Timer {
	return &timedTimer{
		duration:    duration,
		periodic:    false,
		ticker:      nil,
		subscribers: make([]func(), 0),
	}
}

type timedTimer struct {
	duration    time.Duration
	periodic    bool
	ticker      *time.Ticker
	stopChannel chan struct{}
	subscribers []func()
}

func (t *timedTimer) Periodic() Timer {
	t.periodic = true
	return t
}

func (t *timedTimer) Subscribe(subscriber func()) Timer {
	t.subscribers = append(t.subscribers, subscriber)
	return t
}

func (t *timedTimer) Start() {
	t.ticker = time.NewTicker(t.duration)
	t.stopChannel = make(chan struct{})
	go func() {
		for {
			select {
			case <-t.ticker.C:
				for _, subscriber := range t.subscribers {
					subscriber()
				}
				if !t.periodic {
					t.Stop()
				}
			case <-t.stopChannel:
				t.ticker.Stop()
				return
			}
		}
	}()
}

func (t *timedTimer) Stop() {
	t.stopChannel <- struct{}{}
}
