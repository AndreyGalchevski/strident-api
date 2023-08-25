package albums

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

func handleGetAlbums(w http.ResponseWriter, r *http.Request) {
	albums, err := getAlbums()

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, albums)
}

func handleGetAlbumByID(w http.ResponseWriter, r *http.Request) {
	album, err := getAlbumByID(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusNotFound, ErrAlbumNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusOK, album)
}

func handlePostAlbum(w http.ResponseWriter, r *http.Request) {
	var params db.Album

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

	newAlbumID, err := createAlbum(params)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, newAlbumID)
}

func handlePatchAlbum(w http.ResponseWriter, r *http.Request) {
	var params db.Album

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

	ok, err := updateAlbum(mux.Vars(r)["id"], params)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrAlbumNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}

func handleDeleteAlbum(w http.ResponseWriter, r *http.Request) {
	ok, err := deleteAlbum(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrAlbumNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}
