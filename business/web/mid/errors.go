package mid

import (
	"context"
	"net/http"

	"github.com/Loner1024/service/business/sys/validate"

	"github.com/Loner1024/service/foundation/web"
	"go.uber.org/zap"
)

func Errors(log *zap.SugaredLogger) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			v, err := web.GetValues(ctx)
			if err != nil {
				return web.NewShutdownError("web value missing from context")
			}

			if err := handler(ctx, w, r); err != nil {
				log.Errorw("ERROR", "traceid", v.TraceID, "ERROR", err)

				var er validate.ErrorResponse
				var status int
				switch act := validate.Cause(err).(type) {
				case validate.FieldsErrors:
					er = validate.ErrorResponse{
						Error:  "data validation error",
						Fields: act.Error(),
					}
					status = http.StatusBadRequest

				case *validate.RequestError:
					er = validate.ErrorResponse{
						Error: act.Error(),
					}
					status = act.Status

				default:
					er = validate.ErrorResponse{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}

				// Respond with the error back to the client.
				if err := web.Respond(ctx, w, er, status); err != nil {
					return err
				}

				// If we receive the shutdown err we need to return it
				// back to the base handler to shutdown the service.
				if ok := web.IsShutdown(err); ok {
					return err
				}
			}
			return nil
		}

		return h
	}

	return m
}
