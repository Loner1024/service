package testgrp

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type Handlers struct {
	Log *zap.SugaredLogger
}

func (h Handlers) Test(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(response)
	
	statusCode := http.StatusOK
	
	h.Log.Infow("liveness", "statusCode", statusCode, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
}
