package reqid

import "context"

type reqIDCtxKey struct{}

var reqIDContextKey = reqIDCtxKey{}

func FromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	reqID, ok := ctx.Value(reqIDContextKey).(string)
	if !ok {
		return ""
	}
	return reqID
}

func ToContext(ctx context.Context, reqID string) context.Context {
	if ctx == nil {
		return nil
	}
	return context.WithValue(ctx, reqIDContextKey, reqID)
}
