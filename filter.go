package assertions

import (
	"context"
	"fmt"
)

const (
	success                = ""
	needExactValues        = "This assertion requires exactly %d comparison values (you provided %d)."
	needNonEmptyCollection = "This assertion requires at least 1 comparison value (you provided 0)."
	needFewerValues        = "This assertion allows %d or fewer comparison values (you provided %d)."
)

func need(needed int, expected []interface{}) string {
	if len(expected) != needed {
		return fmt.Sprintf(needExactValues, needed, len(expected))
	}
	return success
}

func atLeast(minimum int, expected []interface{}) string {
	if len(expected) < minimum {
		return needNonEmptyCollection
	}
	return success
}

func atMost(max int, expected []interface{}) string {
	if len(expected) > max {
		return fmt.Sprintf(needFewerValues, max, len(expected))
	}
	return success
}

// Context это интерфейс для получения контекста из произвольных объектов
type Context interface {
	// Context возвращает стандартный контекст
	Context() context.Context
}

func getContext(v interface{}) (context.Context, string) {
	if ctx, ok := v.(Context); ok {
		return ctx.Context(), success
	}
	if ctx, ok := v.(context.Context); ok {
		return ctx, success
	}
	return nil, shouldUseContext
}
