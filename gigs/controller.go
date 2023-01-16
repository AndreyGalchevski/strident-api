package gigs

import (
	"net/http"

	"github.com/AndreyGalchevski/strident-api/db"
	"github.com/AndreyGalchevski/strident-api/http_wrapper"
	"github.com/AndreyGalchevski/strident-api/validation"
	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate = validator.New()
var decoder = form.NewDecoder()

func handleGetGigs(w http.ResponseWriter, r *http.Request) {
	gigs, err := getGigs()

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, gigs)
}

func handleGetGigByID(w http.ResponseWriter, r *http.Request) {
	gig, err := getGigByID(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusNotFound, ErrGigNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusOK, gig)
}

func handlePostGig(w http.ResponseWriter, r *http.Request) {
	var params db.Gig

	r.ParseMultipartForm(validation.FormDataLimit)

	err := decoder.Decode(&params, r.Form)

	if err != nil {
		http_wrapper.Failure(w, http.StatusBadRequest, nil)
		return
	}

	err = validate.Struct(&params)

	if err != nil {
		http_wrapper.Failure(w, http.StatusUnprocessableEntity, validation.ErrMissingFields)
		return
	}

	image, _, err := r.FormFile("image")

	if err != nil {
		http_wrapper.Failure(w, http.StatusBadRequest, err)
		return
	}

	if image != nil {
		defer image.Close()
	}

	newGigID, err := createGig(params, image)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, newGigID)
}

func handlePatchGig(w http.ResponseWriter, r *http.Request) {
	var params db.Gig

	r.ParseMultipartForm(validation.FormDataLimit)

	err := decoder.Decode(&params, r.Form)

	if err != nil {
		http_wrapper.Failure(w, http.StatusBadRequest, nil)
		return
	}

	err = validate.Struct(&params)

	if err != nil {
		http_wrapper.Failure(w, http.StatusUnprocessableEntity, validation.ErrMissingFields)
		return
	}

	image, _, _ := r.FormFile("image")

	if image != nil {
		defer image.Close()
	}

	ok, err := updateGig(mux.Vars(r)["id"], params, image)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrGigNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}

func handleDeleteGig(w http.ResponseWriter, r *http.Request) {
	ok, err := deleteGig(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrGigNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}
