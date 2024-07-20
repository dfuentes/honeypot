package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

var port int = 8080

//go:embed warning.sh
var script []byte

func init() {
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
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), &s)
}
