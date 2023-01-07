package auth

import (
	"net/http"

	"github.com/AndreyGalchevski/strident-api/http_wrapper"
)

func VerifyAuthorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(AUTH_COOKIE_NAME)

		if err != nil {
			http_wrapper.Failure(w, http.StatusUnauthorized, nil)
			return
		}

		_, err = VerifyToken(cookie.Value)

		if err != nil {
			http_wrapper.Failure(w, http.StatusUnauthorized, nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}
