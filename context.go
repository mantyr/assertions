package assertions

import (
	"fmt"
	"time"
)

// ShouldBeClosedBefore проверяет что контекст закрылся
// до истечения периода времени
// Поддерживает интерфейс Context и context.Context
func ShouldBeClosedBefore(actual interface{}, expected ...interface{}) string {
	if fail := need(1, expected); fail != success {
		return fail
	}
	ctx, name, fail := getContext(actual)
	if fail != success {
		return fail
	}
	duration, ok := expected[0].(time.Duration)
	if !ok {
		return shouldUseDuration
	}
	select {
	case <-time.After(duration):
		return fmt.Sprintf(shouldClosedBefore, name, duration)
	case <-ctx.Done():
		return ""
	}
	return ""
}

// ShouldNotBeClosedBefore проверяет что контекст не успел закрыться
// до истечения времени
// Поддерживает интерфейс Context и context.Context
func ShouldNotBeClosedBefore(actual interface{}, expected ...interface{}) string {
	if fail := need(1, expected); fail != success {
		return fail
	}
	ctx, name, fail := getContext(actual)
	if fail != success {
		return fail
	}
	duration, ok := expected[0].(time.Duration)
	if !ok {
		return shouldUseDuration
	}
	select {
	case <-time.After(duration):
		return ""
	case <-ctx.Done():
		return fmt.Sprintf(shouldNotClosedBefore, name, duration)
	}
	return ""
}
