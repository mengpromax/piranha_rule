package gray

import "context"

const (
	ConstantKey = "AAA"
)

func IsEnableV3(ctx context.Context, key string) bool {
	return true
}
