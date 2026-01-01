package apiserver

import (
	"net/http"

	"studio-backend/internal/transport/apiserver/gen"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(impl gen.StrictServerInterface) (http.Handler, error) {
	root := chi.NewRouter()

	// TODO MWs:
	// recover from panic
	// healthcheck
	root.Use(middleware.NoCache)
	root.Use(middleware.RealIP)
	root.Use(middleware.CleanPath)
	root.Use(middleware.RequestID)
	// request metadata
	// log
	// trace
	// metrics
	// response metadata
	// auth

	opts := gen.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  func(w http.ResponseWriter, r *http.Request, err error) {},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {},
	}
	handlerImpl := gen.NewStrictHandlerWithOptions(impl, nil, opts)

	return gen.HandlerWithOptions(
		handlerImpl,
		gen.ChiServerOptions{
			BaseRouter:       root,
			ErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {},
		},
	), nil
}
