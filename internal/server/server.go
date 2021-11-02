package server

import (
	"context"
	"net/http"
	"time"

	"tServerOra/internal/handlers"
	"tServerOra/internal/middlewares"
	"tServerOra/internal/models"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	http.Server
}

//Start server with router.
func (s *Server) Start(ctx context.Context, repo models.Repository, opt models.Options) {
	fs := http.FileServer(http.Dir("./html"))
	r := chi.NewRouter()
	handlers.NewHandlers(repo)
	middlewares.NewCookie(repo)

	r.Use(middlewares.SetCookieUser, middlewares.ZipHandlerRead, middlewares.ZipHandlerWrite)
	//r.Use(middlewares.ZipHandlerRead, middlewares.ZipHandlerWrite)
	r.Get("/*", fs.ServeHTTP)
	r.Get("/ping*", handlers.HandlerCheckDBConnect)
	r.Post("/api/savetc*", handlers.HandlerTCPost)
	// r.Delete("/api/user/urls", handlers.HandlerDeleteUserUrls)
	s.Addr = opt.ServAddr()
	s.Handler = r

	go s.ListenAndServe()
	<-ctx.Done()
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	s.Shutdown(ctx)
}
