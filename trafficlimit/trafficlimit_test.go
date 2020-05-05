package trafficlimit

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	redisutil "leetcode/redis"
)

func Test_allowActionCommon(t *testing.T) {
	const (
		userID = "userID"
		action = "login"
	)
	tests := []struct {
		name     string
		duration int
		maxCount int
		expect   []bool
		userID   string
		action   string
	}{
		{
			name:     "duration: 5s maxCount: 5 => 10 allows",
			duration: 5,
			maxCount: 5,
			expect: []bool{true, true, true, true, true,
				true, true, true, true, true},
		},
		{
			name:     "duration: 2s maxCount:1 => 5 allows",
			duration: 2,
			maxCount: 1,
			expect:   []bool{true, true, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				rc, err := redisutil.NewRedisClient()
				if err != nil {
					t.Fatal(err)
				}
				rc.ZRemRangeByScore(fmt.Sprintf("zset:%s:%s", userID, action), "-1", "9999999999999999")
			})

			got := make([]bool, 0)
			startTime := time.Now()
			for time.Since(startTime).Seconds() < 10 {
				isAllow, err := allowActionCommon(userID, action, tt.duration, tt.maxCount)
				if err != nil {
					t.Fatal(err)
				}
				// log.Print("isAllow:::::", isAllow)
				if isAllow {
					got = append(got, true)
					// Sleep a while.
					time.Sleep(10 * time.Millisecond)
					continue
				}

				time.Sleep(time.Duration(tt.duration)*time.Second + 10*time.Millisecond)
			}
			assert.Equal(t, tt.expect, got)
		})
	}
}

func Test_allowActionFunnel(t *testing.T) {
	tests := []struct {
		name     string
		cap      float32
		duration float32
		maxCount float32
		expect   []bool
	}{
		{
			name:     "cap=5 duration=5s maxCount=5 => speed 1/s",
			cap:      5,
			duration: 5,
			maxCount: 5,
			expect: []bool{true, true, true, true, true,
				true, true, true, true, true},
		},
		{
			name:     "cap=3 duration=2s maxCount=1 => speed 0.5/s",
			cap:      3,
			duration: 2,
			maxCount: 1,
			expect: []bool{true, true, true,
				true, true, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			funnel := newFunnel(tt.cap, tt.maxCount/tt.duration)
			startTime := time.Now()
			got := make([]bool, 0)
			for time.Since(startTime).Seconds() < 10 {
				if isAllow := allowAction(funnel); isAllow {
					got = append(got, true)
					// log.Print("isAllow:::::", isAllow)
					// Sleep a while.
					time.Sleep(1 * time.Millisecond)
					continue
				}
				time.Sleep(time.Duration(tt.duration) * time.Second)
			}

			assert.Equal(t, tt.expect, got)
		})
	}
}
