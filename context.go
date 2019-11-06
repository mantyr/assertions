package assertions

import (
	"context"
	"fmt"
	"time"
)

// ShouldBeClosedBefore проверяет что контекст закрылся
// до истечения периода времени
func ShouldBeClosedBefore(actual interface{}, expected ...interface{}) string {
	if fail := need(1, expected); fail != success {
		return fail
	}
	ctx, ok := actual.(context.Context)
	if !ok {
		return shouldUseContext
	}
	duration, ok := expected[0].(time.Duration)
	if !ok {
		return shouldUseDuration
	}
	select {
	case <-time.After(duration):
		return fmt.Sprintf(shouldClosedBefore, duration)
	case <-ctx.Done():
		return ""
	}
	return ""
}

// ShouldNotBeClosedBefore проверяет что контекст не успел закрыться
// до истечения времени
func ShouldNotBeClosedBefore(actual interface{}, expected ...interface{}) string {
	if fail := need(1, expected); fail != success {
		return fail
	}
	ctx, ok := actual.(context.Context)
	if !ok {
		return shouldUseContext
	}
	duration, ok := expected[0].(time.Duration)
	if !ok {
		return shouldUseDuration
	}
	select {
	case <-time.After(duration):
		return ""
	case <-ctx.Done():
		return fmt.Sprintf(shouldNotClosedBefore, duration)
	}
	return ""
}
