package testgrp

import (
	"context"
	"net/http"

	"github.com/Loner1024/service/foundation/web"

	"go.uber.org/zap"
)

type Handlers struct {
	Log *zap.SugaredLogger
}

func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Response(ctx, w, status, http.StatusOK)
}
