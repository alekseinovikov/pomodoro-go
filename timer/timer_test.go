package timer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockStub struct {
	calledCount int
}

func newMockStub() *mockStub {
	return &mockStub{calledCount: 0}
}

func (m *mockStub) OnTimer() {
	m.calledCount++
}

type timerTest struct {
	timerDuration time.Duration
	waitDuration  time.Duration
	result        int
}

var oneTimeTimerTests = []timerTest{
	{1 * time.Millisecond, 10 * time.Millisecond, 1},
	{10 * time.Millisecond, 1 * time.Millisecond, 0},
}

var periodicTimerTests = []timerTest{
	{10 * time.Millisecond, 25 * time.Millisecond, 2},
	{10 * time.Millisecond, 15 * time.Millisecond, 1},
}

func TestOneTimeTimer(t *testing.T) {
	for _, tt := range oneTimeTimerTests {
		timer := NewTimer(tt.timerDuration)
		mockStub := newMockStub()
		timer.Subscribe(mockStub.OnTimer)
		timer.Start()
		time.Sleep(tt.waitDuration)
		assert.Equal(t, tt.result, mockStub.calledCount)
	}
}

func TestPeriodicTimer(t *testing.T) {
	for _, tt := range periodicTimerTests {
		timer := NewTimer(tt.timerDuration).Periodic()
		mockStub := newMockStub()
		timer.
			Subscribe(mockStub.OnTimer).
			Start()
		time.Sleep(tt.waitDuration)
		timer.Stop()
		assert.Equal(t, tt.result, mockStub.calledCount)
	}
}
