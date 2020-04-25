package main

import (
	"sync"
	"time"
)

type funnel struct {
	mux sync.Mutex

	cap           float32
	left          float32
	speed         float32
	lastTimeStamp time.Time
}

func newFunnel(cap, speed float32) *funnel {
	return &funnel{
		cap:           cap,
		left:          cap,
		speed:         speed,
		lastTimeStamp: time.Now(),
	}
}

func (f *funnel) makeSpace() {
	now := time.Now()
	outflow := float32(now.Sub(f.lastTimeStamp).Seconds()) * f.speed

	if f.left+outflow > f.cap {
		f.left = f.cap
		f.lastTimeStamp = now
		return
	}
	f.left += outflow
	f.lastTimeStamp = now
}

func (f *funnel) water(inflow float32) bool {
	if inflow > f.left {
		return false
	}
	f.left -= inflow
	return true
}

// TODO: should get from redis.
func allowAction(f *funnel) bool {
	f.mux.Lock()
	defer f.mux.Unlock()

	f.makeSpace()
	return f.water(1)
}
