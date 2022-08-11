package server

import (
	"net/http"
	"telegrambot/package/repository"
	"telegrambot/utils/logrus_log"

	"github.com/sirupsen/logrus"
)

type Server struct {
	server *http.Server
	logrus *logrus_log.Logger
	repo   *repository.Repository
}

func NewServer(repo *repository.Repository, logrus *logrus_log.Logger) *Server {
	return &Server{
		repo:   repo,
		logrus: logrus,
	}
}

func (s *Server) Start() error {
	s.server = &http.Server{
		Handler: s,
		Addr:    ":8080",
	}
	logrus.Info("Start Server")
	return s.server.ListenAndServe()
}
