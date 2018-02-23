package util

import (
	"context"
	"io"
	"testing"
	"time"
)

func TestRetryForever(t *testing.T) {

	var err error

	expiredAt := time.Now().Add(2 * time.Second)
	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	RetryForever(ctx, "test1", func() error {
		if time.Now().Before(expiredAt) {
			err = io.EOF
			return io.EOF
		}
		err = nil
		return nil
	}, 400*time.Millisecond,
	)

	if err == nil {
		t.Errorf("unexpected nil error")
	}

	expiredAt = time.Now().Add(500 * time.Millisecond)
	ctx, _ = context.WithTimeout(context.Background(), time.Second)
	RetryForever(ctx, "test2", func() error {
		if time.Now().Before(expiredAt) {
			err = io.EOF
			return io.EOF
		}
		err = nil
		return nil
	}, 200*time.Millisecond,
	)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

}
