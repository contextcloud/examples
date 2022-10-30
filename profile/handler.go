package function

import (
	"context"
	"net/http"

	"github.com/contextcloud/graceful/config"
)

func NewHandler(ctx context.Context, cfg *config.Config) (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}), nil
}
