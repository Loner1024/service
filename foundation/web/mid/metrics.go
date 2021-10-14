package mid

import (
	"context"
	"net/http"

	"github.com/Loner1024/service/business/sys/metrics"

	"github.com/Loner1024/service/foundation/web"
)

func Metrics() web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			ctx = metrics.Set(ctx)

			err := handler(ctx, w, r)

			metrics.AddRequest(ctx)
			metrics.AddGoroutines(ctx)
			if err != nil {
				metrics.AddErrors(ctx)
			}
			return err
		}
		return h
	}
	return m
}
