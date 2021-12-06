package handlers

import (
	"expvar"
	"github.com/Loner1024/service/app/services/sales-api/handlers/debug/checkgrp"
	"github.com/Loner1024/service/app/services/sales-api/handlers/v1/testgrp"
	"github.com/dimfeld/httptreemux/v5"
	"go.uber.org/zap"
	"net/http"
	"net/http/pprof"
	"os"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

func APIMux(cfg APIMuxConfig) *httptreemux.ContextMux {
	mux := httptreemux.NewContextMux()
	
	tgh := testgrp.Handlers{
		Log: cfg.Log,
	}
	
	mux.Handle(http.MethodGet, "/v1/test", tgh.Test)
	
	return mux
}

func DebugStandardLibraryMux() *http.ServeMux {
	mux := http.NewServeMux()
	
	// Register all the standard library debug endpoints.
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())
	
	return mux
}

func DebugMux(build string, log *zap.SugaredLogger) http.Handler {
	mux := DebugStandardLibraryMux()
	
	cgh := checkgrp.Handlers{
		Build: build,
		Log:   log,
	}
	
	mux.HandleFunc("/debug/readiness", cgh.Readiness)
	mux.HandleFunc("/debug/liveness", cgh.Liveness)
	
	return mux
}
