package websockify

import (
	"net/http"
	"testing"
)

func TestWebsock01(t *testing.T) {
	logger.Printf("listening on %s\n", cfg.Server.Addr())
	http.HandleFunc("/", handleConnection)
	http.ListenAndServe(cfg.Server.Addr(), nil)
}
