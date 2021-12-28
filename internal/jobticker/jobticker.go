package jobticker

import (
	"context"
	"time"

	"github.com/SayYoungMan/gitmuzik-backend/internal/logger"
)

const (
	INTERVAL_PERIOD time.Duration = 24 * time.Hour
	HOUR_TO_TICK    int           = 19 // Makes it to tick at 19:00 everyday
)

type jobTicker struct {
	Timer *time.Timer
}

func getNextTickDuration(ctx context.Context) time.Duration {
	now := time.Now()
	// Set the next tick at HOUR_TO_TICK today
	nextTick := time.Date(now.Year(), now.Month(), now.Day(), HOUR_TO_TICK, 0, 0, 0, time.Local)
	if nextTick.Before(now) {
		// If it is already past HOUR_TO_TICK, set it to next day
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	logger.FromContext(ctx).Info("Next tick is scheduled at: ", nextTick)
	// Return the duration from now to next tick
	return time.Until(nextTick)
}

func NewJobTicker(ctx context.Context) jobTicker {
	return jobTicker{
		time.NewTimer(getNextTickDuration(ctx)),
	}
}

func (jt jobTicker) UpdateJobTicker(ctx context.Context) {
	jt.Timer.Reset(getNextTickDuration(ctx))
	logger.FromContext(ctx).Info("Successfully updated Job Ticker")
}
