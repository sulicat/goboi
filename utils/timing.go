package utils

import "time"

type WaitTimer struct {
	time_started  float64
	time_reset    float64
	duratation_ms float64
}

func CreateWaitTimer(duration_s float64) WaitTimer {
	out := WaitTimer{}

	out.time_started = float64(time.Now().UnixMilli())
	out.duratation_ms = duration_s * 1000
	out.Reset()
	return out
}

func (t *WaitTimer) Reset() {
	t.time_reset = float64(time.Now().UnixMilli())
}

func (t *WaitTimer) SetDuration(d float64) {
	t.duratation_ms = d
}

func (t *WaitTimer) Check() bool {
	now := float64(time.Now().UnixMilli())
	return (now - t.time_reset) > t.duratation_ms
}
