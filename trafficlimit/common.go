package trafficlimit

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v7"

	redisutil "leetcode/redis"
)

func allowActionCommon(userID, action string, durationSec, count int) (bool, error) {
	rc, err := redisutil.NewRedisClient()
	if err != nil {
		return false, err
	}
	var (
		format           = "20060102150405.000"
		key              = fmt.Sprintf("zset:%s:%s", userID, action)
		now              = time.Now().Local()
		deleteScore      = now.Add(time.Duration(-2*durationSec) * time.Second).Format(format)
		beforeScore      = now.Add(time.Duration(-1*durationSec) * time.Second).Format(format)
		nowScore         = now.Format(format)
		nowScoreFloat, _ = strconv.ParseFloat(nowScore, 64)
	)

	pipe := rc.Pipeline()
	pipe.ZRemRangeByScore(key, "-1", deleteScore)
	pipe.ZAdd(key, &redis.Z{
		Score:  nowScoreFloat,
		Member: nowScoreFloat,
	})
	pipe.ZCount(key, beforeScore, nowScore)

	cmds, err := pipe.ExecContext(context.Background())
	if err != nil {
		return false, err
	}
	c, ok := cmds[len(cmds)-1].(*redis.IntCmd)
	if !ok {
		return false, errors.New("convert fail")
	}
	val, _ := c.Result()
	if val > int64(count) {
		return false, nil
	}
	return true, nil
}

type funnel struct {
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

func allowAction(f *funnel) bool {
	f.makeSpace()
	return f.water(1)
}
