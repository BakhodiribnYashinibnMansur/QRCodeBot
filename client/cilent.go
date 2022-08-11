package client

import (
	"net/http"
	"telegrambot/config"
	"telegrambot/package/repository"
	"telegrambot/utils/logrus_log"
	"time"

	"github.com/sirupsen/logrus"
)

type Client struct {
	config *config.Configuration
	client *http.Client
	logrus *logrus_log.Logger
	repo   *repository.Repository
}

func NewClient(config *config.Configuration, repo *repository.Repository, logrus *logrus_log.Logger) *Client {
	return &Client{
		repo:   repo,
		logrus: logrus,
	}
}

func (clnt *Client) Start() *Client {
	client := &http.Client{Timeout: time.Duration(4) * time.Second}
	logrus.Info("Start Client")
	return &Client{
		client: client,
	}
}
