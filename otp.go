package main

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type OTP struct {
	Key     string
	Created time.Time
}

type RetentionMap map[string]OTP

// NewOTP creates and adds a new otp to the map
func (rm RetentionMap) NewOTP() OTP {
	o := OTP{
		Key:     uuid.NewString(),
		Created: time.Now(),
	}

	rm[o.Key] = o
	return o
}

// VerifyOTP will make sure an OTP exists and return true if so
// It will also delete the key, so it can't be reused
func (rm RetentionMap) VerifyOTP(otp string) bool {
	// Verify OTP is existing
	if _, ok := rm[otp]; !ok {
		// otp does not exist
		return false
	}
	delete(rm, otp)
	return true
}

// Retention will make sure old OTPs are removed
// Is Blocking, so run as a Goroutine
func (rm RetentionMap) Retention(ctx context.Context, retentionPeriod time.Duration) {
	ticker := time.NewTicker(400 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			for _, otp := range rm {
				// Add Retention to Created and check if it is expired
				if otp.Created.Add(retentionPeriod).Before(time.Now()) {
					delete(rm, otp.Key)
				}
			}
		case <-ctx.Done():
			return
		}
	}
}

// NewRetentionMap will create a new retentionmap and start the retention given the set period
func NewRetentionMap(ctx context.Context, retentionPeriod time.Duration) RetentionMap {
	rm := make(RetentionMap)
	go rm.Retention(ctx, retentionPeriod)
	return rm
}
