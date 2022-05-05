package service

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/keloran/go-healthcheck"
	"github.com/keloran/go-probe"
	"net/http"

	bugLog "github.com/bugfixes/go-bugfixes/logs"
	bugMiddleware "github.com/bugfixes/go-bugfixes/middleware"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	"github.com/retro-board/company-service/internal/config"
	"github.com/retro-board/company-service/internal/user"
)

type Service struct {
	Config *config.Config
}

func (s *Service) Start() error {
	bugLog.Local().Info("Starting User")

	logger := httplog.NewLogger("user-api", httplog.Options{
		JSON: true,
	})

	allowedOrigins := []string{
		"http://localhost:8080",
		"https://retro-board.it",
		"https://*.retro-board.it",
	}
	if s.Config.Development {
		allowedOrigins = append(allowedOrigins, "http://*")
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-User-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/ping"))

	r.Route("/", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(c.Handler)
		r.Use(bugMiddleware.BugFixes)
		r.Use(httplog.RequestLogger(logger))

		r.Get("/login", user.NewUser(s.Config).LoginHandler)
		r.Get("/callback", user.NewUser(s.Config).CallbackHandler)
		r.Post("/verify", user.NewUser(s.Config).VerifyHandler)
		r.Get("/permission", user.NewUser(s.Config).PermissionHandler)
	})

	r.Get("/health", healthcheck.HTTP)
	r.Get("/probe", probe.HTTP)

	bugLog.Local().Infof("listening on %d\n", s.Config.Local.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.Config.Local.Port), r); err != nil {
		return bugLog.Errorf("port failed: %+v", err)
	}

	return nil
}
