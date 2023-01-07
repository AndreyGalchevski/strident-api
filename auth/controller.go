package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/AndreyGalchevski/strident-api/http_wrapper"
	"github.com/AndreyGalchevski/strident-api/validation"
	"github.com/go-playground/validator/v10"
)

const AUTH_COOKIE_NAME = "stridentToken"

type Credentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

var validate = validator.New()

func handlePostLogin(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		http_wrapper.Failure(w, http.StatusBadRequest, errors.New("please try again"))
		return
	}

	err = validate.Struct(&credentials)

	if err != nil {
		http_wrapper.Failure(w, http.StatusUnprocessableEntity, validation.ErrMissingFields)
		return
	}

	token, err := login(credentials)

	if err != nil {
		http_wrapper.Failure(w, http.StatusUnauthorized, err)
		return
	}

	isProd := os.Getenv("APP_ENV") == "prod"

	domain := ""

	if isProd {
		domain = strings.TrimSuffix(os.Getenv("WEB_APP_URL"), "/")
	}

	cookie := http.Cookie{
		Name:     AUTH_COOKIE_NAME,
		Value:    token,
		Path:     "/",
		MaxAge:   int(TokenMaxAge.Seconds()),
		HttpOnly: true,
		Secure:   isProd,
		SameSite: http.SameSiteLaxMode,
		Domain:   domain,
	}

	http.SetCookie(w, &cookie)

	http_wrapper.Success(w, http.StatusNoContent, nil)
}

func handleGetVerify(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie(AUTH_COOKIE_NAME)

	if err != nil {
		http_wrapper.Failure(w, http.StatusUnauthorized, errors.New("session expired"))
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}
