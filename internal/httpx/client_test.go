package httpx_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LYH2263/go-fetch-timeout/internal/httpx"
)

func TestOK(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	defer srv.Close()
	b, err := httpx.New(srv.Client()).Get(context.Background(), srv.URL)
	if err != nil || string(b) != "ok" {
		t.Fatalf("%s %v", b, err)
	}
}
