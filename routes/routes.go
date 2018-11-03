package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rickbassham/goapi/router"
)

type HandlerService interface {
	Foo(w http.ResponseWriter, r *http.Request)
}

type API struct {
	handler HandlerService
}

func NewAPI(handler HandlerService) router.RouteCreater {
	return &API{
		handler: handler,
	}
}

func (svc *API) CreateRoutes(r chi.Router) chi.Router {
	r.Get("/foo", svc.handler.Foo)

	return r
}
