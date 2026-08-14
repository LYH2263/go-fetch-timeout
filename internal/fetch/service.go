package fetch

import (
	"context"
	"net/http"

	"github.com/LYH2263/go-fetch-timeout/internal/httpx"
	"github.com/LYH2263/go-fetch-timeout/internal/model"
)

type Service struct{ C *httpx.Client }

func New(hc *http.Client) *Service { return &Service{C: httpx.New(hc)} }

func (s *Service) Get(ctx context.Context, url string) (model.Result, error) {
	b, err := s.C.Get(ctx, url)
	if err != nil {
		return model.Result{}, err
	}
	return model.Result{URL: url, Body: b}, nil
}
