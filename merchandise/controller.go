package merchandise

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

func handleGetMerchandise(w http.ResponseWriter, r *http.Request) {
	merchandise, err := getMerchandise()

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, merchandise)
}

func handleGetMerchandiseByID(w http.ResponseWriter, r *http.Request) {
	merchandise, err := getMerchandiseByID(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusNotFound, ErrMerchandiseNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusOK, merchandise)
}

func handlePostMerchandise(w http.ResponseWriter, r *http.Request) {
	var params db.Merchandise

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

	newMerchandiseID, err := createMerchandise(params, image)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, newMerchandiseID)
}

func handlePatchMerchandise(w http.ResponseWriter, r *http.Request) {
	var params db.Merchandise

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

	ok, err := updateMerchandise(mux.Vars(r)["id"], params, image)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrMerchandiseNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}

func handleDeleteMerchandise(w http.ResponseWriter, r *http.Request) {
	ok, err := deleteMerchandise(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrMerchandiseNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}
