package timer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockStub struct {
	called bool
}

func newMockStub() *mockStub {
	return &mockStub{called: false}
}

func (m *mockStub) OnTimer() {
	m.called = true
}

type timerTest struct {
	timerDuration time.Duration
	waitDuration  time.Duration
	result        bool
}

var tests = []timerTest{
	{1 * time.Millisecond, 10 * time.Millisecond, true},
	{10 * time.Millisecond, 1 * time.Millisecond, false},
}

func TestTimer(t *testing.T) {
	for _, tt := range tests {
		timer := NewTimer(tt.timerDuration)
		mockStub := newMockStub()
		timer.Subscribe(mockStub)
		timer.Start()
		time.Sleep(tt.waitDuration)
		assert.Equal(t, tt.result, mockStub.called)
	}
}
