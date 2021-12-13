package mid

import (
	"context"
	"github.com/Loner1024/service/foundation/web"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Logger(log *zap.SugaredLogger) web.Middleware {
	return func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			
			v, err := web.GetValues(ctx)
			if err != nil {
				return err
			}
			
			log.Infow("request started", "traceid", v.TraceID, "method", r.Method, "path", r.URL.Path,
				"remoteaddr", r.RemoteAddr)
			
			err = handler(ctx, w, r)
			
			log.Infow("request started", "traceid", v.TraceID, "method", r.Method, "path",
				r.URL.Path, "remoteaddr", r.RemoteAddr, "statuscode", v.StatusCode, "since", time.Since(v.Now))
			
			return err
		}
		return h
	}
}
