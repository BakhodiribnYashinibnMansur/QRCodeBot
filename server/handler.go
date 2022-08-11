package server

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.logrus.Debug("received unavailable HTTP method request",
			r.Method)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	chatIDQuery := r.URL.Query().Get("chat_id")
	if chatIDQuery == "" {
		s.logrus.Debug("received empty chat_id query param")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatID, err := strconv.ParseInt(chatIDQuery, 10, 64)
	if err != nil {
		s.logrus.Debug("received invalid chat_id query param",
			chatIDQuery)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	logrus.Info("received chat_id", chatID)
	w.Header().Set("Location", "chatID")

	w.WriteHeader(http.StatusMovedPermanently)
}

// func (s *Service) createAccessToken(ctx context.Context, chatID int64) error {
// 	requestToken, err := s.storage.Get(chatID, storage.RequestTokens)
// 	if err != nil {
// 		return errors.WithMessage(err, "failed to get request token")
// 	}

// 	authResp, err := s.client.Authorize(ctx, requestToken)
// 	if err != nil {
// 		return errors.WithMessage(err, "failed to authorize at Pocket")
// 	}

// 	if err := s.storage.Save(chatID, authResp.AccessToken, storage.AccessTokens); err != nil {
// 		return errors.WithMessage(err, "failed to save access token to storage")
// 	}

// 	return nil
// }
