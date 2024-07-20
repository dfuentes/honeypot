package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

var port int = 8080

//go:embed warning.sh
var script []byte

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if envPort, ok := os.LookupEnv("PORT"); ok {
		envPortI, err := strconv.Atoi(envPort)
		if err != nil {
			slog.Error("failed to parse port env var")
			os.Exit(1)
		}
		port = envPortI
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.InfoContext(ctx, "request", "url", r.URL.String(),
		"remote", r.RemoteAddr,
		"ua", r.Header.Get(http.CanonicalHeaderKey("user-agent")),
		"xff", r.Header.Get(http.CanonicalHeaderKey("x-forwarded-for")),
	)
	w.Write(script)
}

func main() {
	s := http.ServeMux{}
	s.HandleFunc("GET /", handle)

	slog.Info("starting server", "port", port)

	server := &http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%d", port),
		Handler:        &s,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server error", "err", err)
		os.Exit(1)
	}
}
