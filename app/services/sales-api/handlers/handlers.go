package handlers

import (
	"expvar"
	"github.com/Loner1024/service/app/services/sales-api/handlers/debug/checkgrp"
	"github.com/Loner1024/service/app/services/sales-api/handlers/v1/testgrp"
	"github.com/Loner1024/service/business/web/mid"
	"github.com/Loner1024/service/foundation/web"
	"go.uber.org/zap"
	"net/http"
	"net/http/pprof"
	"os"
)

type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *zap.SugaredLogger
}

func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(
		cfg.Shutdown,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Metrics(),
		mid.Panics(),
	)
	v1(app, cfg)
	
	return app
}

func v1(app *web.App, cfg APIMuxConfig) {
	const version = "v1"
	tgh := testgrp.Handlers{
		Log: cfg.Log,
	}
	app.Handle(http.MethodGet, version, "/test", tgh.Test)
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
