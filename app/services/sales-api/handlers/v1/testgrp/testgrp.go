package testgrp

import (
	"context"
	"github.com/Loner1024/service/business/sys/validate"
	"math/rand"
	"net/http"

	"github.com/Loner1024/service/foundation/web"

	"errors"
	"go.uber.org/zap"
)

type Handlers struct {
	Log *zap.SugaredLogger
}

func (h Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100);n%2==0{
		return validate.NewRequestError(errors.New("trust error"),http.StatusBadRequest)
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
