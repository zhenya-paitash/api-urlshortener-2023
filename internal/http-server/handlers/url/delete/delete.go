package delete

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	resp "github.com/zhenya-paitash/api-urlshortener-2023/internal/lib/api/response"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/lib/logger/sl"
	"github.com/zhenya-paitash/api-urlshortener-2023/internal/storage"
)

// go:generate go run github.com/vektra/mockery/v2@v2.38.0 --name=URLDelete
type URLDelete interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, urlDelete URLDelete) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.delete.New"

		log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")
			render.JSON(w, r, resp.Error("invalid request"))
			return
		}

		err := urlDelete.DeleteURL(alias)
		if err != nil {
			if errors.Is(err, storage.ErrURLNotFound) {
				log.Info("url not found", "alias", alias)
				render.JSON(w, r, resp.Error("not found"))
				return
			}
			log.Error("failed to get url", sl.Err(err))
			render.JSON(w, r, resp.Error("internal error"))
			return
		}

    log.Info("alias has been deleted", slog.String("alias", alias))

    render.JSON(w, r, resp.OK())
	}
}
