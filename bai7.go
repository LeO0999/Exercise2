package main

import (
	"context"
	"time"
)

func X(ctx context.Context, k string) int64 {

	before := ctx.Value(k).(int64)
	var now int64
	select {
	case <-time.After(3 * time.Second):
		now = time.Now().UnixNano()

	}

	return now - before

}
