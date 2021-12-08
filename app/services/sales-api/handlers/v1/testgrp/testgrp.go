package testgrp

import (
	"context"
	"github.com/Loner1024/service/foundation/web"
	"go.uber.org/zap"
	"net/http"
)

type Handlers struct {
	Log *zap.SugaredLogger
}

func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	response := struct {
		Status string
	}{
		Status: "OK",
	}
	
	statusCode := http.StatusOK
	
	h.Log.Infow("liveness", "statusCode", statusCode, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
	
	return web.Respond(ctx, w, response, statusCode)
}
