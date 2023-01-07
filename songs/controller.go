package songs

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

func handleGetSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := getSongs()

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, songs)
}

func handleGetSongByID(w http.ResponseWriter, r *http.Request) {
	song, err := getSongByID(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusNotFound, ErrSongNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusOK, song)
}

func handlePostSong(w http.ResponseWriter, r *http.Request) {
	var params db.Song

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

	newSongID, err := createSong(params)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, newSongID)
}

func handlePatchSong(w http.ResponseWriter, r *http.Request) {
	var params db.Song

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

	ok, err := updateSong(mux.Vars(r)["id"], params)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrSongNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}

func handleDeleteSong(w http.ResponseWriter, r *http.Request) {
	ok, err := deleteSong(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrSongNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}
