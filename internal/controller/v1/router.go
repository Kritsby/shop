package v1

import (
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
	"net/http"
	"shop/internal/service"
)

type Handler struct {
	services service.Stocker
}

func New(services service.Stocker) *Handler {
	return &Handler{services: services}
}

type JSONServer struct{}

func (h *Handler) InitRouter() *bunrouter.Router {
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
	)

	swagHandler := httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)
	bswag := bunrouter.HTTPHandlerFunc(swagHandler)
	router.GET("/swagger/:*", bswag)

	router.WithGroup("/v1", func(g *bunrouter.Group) {
		g.POST("/reserve", h.reserve)
		g.POST("/release", h.reserveRelease)
		g.GET("/amount/:storage", h.amount)
	})

	return router
}

func (h *Handler) responseJSON(w http.ResponseWriter, req bunrouter.Request, code int, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if code != http.StatusOK {
		log.Warn().Msgf("route: %s, http code: %d, error: %v", req.Route(), code, value)
		return bunrouter.JSON(w, bunrouter.H{
			"route":  req.Route(),
			"params": req.Params().Map(),
			"error":  value,
		})
	}

	return bunrouter.JSON(w, bunrouter.H{
		"route":  req.Route(),
		"params": req.Params().Map(),
		"data":   value,
	})
}
