package fetch_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/LYH2263/go-fetch-timeout/internal/fetch"
)

func TestParentCancelStopsRequest(t *testing.T) {
	block := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-block
	}))
	defer srv.Close()
	defer close(block)

	ctx, cancel := context.WithCancel(context.Background())
	s := fetch.New(srv.Client())
	errCh := make(chan error, 1)
	go func() {
		_, err := s.Get(ctx, srv.URL)
		errCh <- err
	}()
	time.Sleep(20 * time.Millisecond)
	cancel()
	select {
	case err := <-errCh:
		if err == nil {
			t.Fatal("want cancel error")
		}
	case <-time.After(2 * time.Second):
		t.Fatal("request not cancelled")
	}
}
