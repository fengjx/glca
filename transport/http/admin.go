package http

import (
	"net/http"
	"strconv"

	"github.com/fengjx/luchen"
	"go.uber.org/zap"

	"github.com/fengjx/glca/connom/auth"
	"github.com/fengjx/glca/current"
)

const (
	requestHeaderDebugUID = "X-Debug-UID"
	requestHeaderToken    = "X-Token"
)

func adminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := luchen.Logger(r.Context())
		var uid int64
		token := r.Header.Get(requestHeaderToken)
		if len(token) > 0 {
			payload, err := auth.Parse(token)
			if err != nil {
				uid = payload.UID
			}
		}
		if uid == 0 && !luchen.IsProd() {
			uidHeader := r.Header.Get(requestHeaderDebugUID)
			if len(uidHeader) > 0 {
				debugUID, err := strconv.ParseInt(uidHeader, 10, 64)
				if err == nil {
					uid = debugUID
					log.Info("set debug uid", zap.Int64("uid", uid))
				}
			}
		}
		if uid > 0 {
			ctx := luchen.WithLogger(r.Context(), log.With(zap.Int64("uid", uid)))
			ctx = current.WithUID(ctx, uid)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}
