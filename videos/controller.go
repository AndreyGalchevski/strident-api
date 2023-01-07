package videos

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

func handleGetVideos(w http.ResponseWriter, r *http.Request) {
	videos, err := getVideos()

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, videos)
}

func handleGetVideoByID(w http.ResponseWriter, r *http.Request) {
	video, err := getVideoByID(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusNotFound, ErrVideoNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusOK, video)
}

func handlePostVideo(w http.ResponseWriter, r *http.Request) {
	var params db.Video

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

	newVideoID, err := createVideo(params)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, newVideoID)
}

func handlePatchVideo(w http.ResponseWriter, r *http.Request) {
	var params db.Video

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

	ok, err := updateVideo(mux.Vars(r)["id"], params)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrVideoNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}

func handleDeleteVideo(w http.ResponseWriter, r *http.Request) {
	ok, err := deleteVideo(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrVideoNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}
