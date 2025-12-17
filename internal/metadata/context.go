package metadata

import "context"

type ctxKey struct{}

var metaCtxKey ctxKey

func ToContext(ctx context.Context, meta RequestMetadata) context.Context {
	return context.WithValue(ctx, metaCtxKey, meta)
}

func FromContext(ctx context.Context) (RequestMetadata, bool) {
	if meta, ok := ctx.Value(metaCtxKey).(RequestMetadata); ok {
		return meta, true
	}
	return RequestMetadata{}, false
}
