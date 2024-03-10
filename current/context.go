package current

import "context"

type (
	uidCtxKey struct{}

	requestActionCtxKey struct{}

	protocolCtxKey struct{}
	methodCtxKey   struct{}
)

func UID(ctx context.Context) int64 {
	val := ctx.Value(uidCtxKey{})
	if val == nil {
		return 0
	}
	if uid, ok := val.(int64); ok {
		return uid
	}
	return 0
}

func WithUID(ctx context.Context, uid int64) context.Context {
	return context.WithValue(ctx, uidCtxKey{}, uid)
}

func RequestAction(ctx context.Context) string {
	val, ok := ctx.Value(requestActionCtxKey{}).(string)
	if !ok {
		return ""
	}
	return val
}

func WithRequestAction(ctx context.Context, action string) context.Context {
	return context.WithValue(ctx, requestActionCtxKey{}, action)
}

// Protocol 请求协议
func Protocol(ctx context.Context) string {
	val, ok := ctx.Value(protocolCtxKey{}).(string)
	if !ok {
		return ""
	}
	return val
}

func WithProtocol(ctx context.Context, protocol string) context.Context {
	return context.WithValue(ctx, protocolCtxKey{}, protocol)
}

// Method 请求协议
func Method(ctx context.Context) string {
	val, ok := ctx.Value(methodCtxKey{}).(string)
	if !ok {
		return ""
	}
	return val
}

func WithMethod(ctx context.Context, method string) context.Context {
	return context.WithValue(ctx, methodCtxKey{}, method)
}
