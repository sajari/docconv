package debug

import (
	"context"
)

type debugEnabledKey struct{}

func debugEnabledInContext(ctx context.Context) bool {
	enabled, ok := ctx.Value(debugEnabledKey{}).(bool)
	if !ok {
		return false
	}
	return enabled
}
