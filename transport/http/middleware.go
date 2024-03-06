package http

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fengjx/luchen"
	"go.uber.org/zap"

	"github.com/fengjx/glca/connom/auth"
	"github.com/fengjx/glca/connom/errno"
	"github.com/fengjx/glca/current"
)

var (
	noAuthPaths = []string{
		"/api/login",
		"/static",
	}
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := luchen.Logger(r.Context())
		w.Header().Set(ResponseHeaderServer, "glca")
		var uid int64
		token := r.Header.Get(RequestHeaderToken)
		if len(token) > 0 {
			payload, expiresAt, err := auth.Parse(token)
			if err != nil {
				log.Warn("parse token err", zap.String("token", token), zap.Error(err))
			}
			uid = payload.UID
			if time.Unix(expiresAt, 0).Sub(time.Now()) < (time.Hour * 24 * 6) {
				refreshToken, _ := auth.GenToken(payload)
				w.Header().Set(ResponseHeaderRefreshToken, refreshToken)
			}
		}
		if uid == 0 && !luchen.IsProd() {
			uidHeader := r.Header.Get(RequestHeaderDebugUID)
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

func adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isNoAuthPath(r) {
			next.ServeHTTP(w, r)
			return
		}
		log := luchen.Logger(r.Context())
		uid := current.UID(r.Context())
		if uid == 0 {
			log.Warn("request unauthorized", zap.String("path", r.URL.Path))
			err := errno.UnauthorizedErr
			WriteData(
				r.Context(),
				w,
				err.HTTPCode,
				&result{
					Status: err.Code,
					Msg:    err.Msg,
				},
			)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isNoAuthPath(r *http.Request) bool {
	p := r.URL.Path
	for _, prefix := range noAuthPaths {
		if strings.HasPrefix(p, prefix) {
			return true
		}
	}
	return false
}