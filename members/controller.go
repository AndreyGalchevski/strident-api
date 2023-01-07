package members

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

func handleGetMembers(w http.ResponseWriter, r *http.Request) {
	members, err := getMembers()

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, members)
}

func handleGetMemberByID(w http.ResponseWriter, r *http.Request) {
	member, err := getMemberByID(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusNotFound, ErrMemberNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusOK, member)
}

func handlePostMember(w http.ResponseWriter, r *http.Request) {
	var params db.Member

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

	newMemberID, err := createMember(params, image)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	http_wrapper.Success(w, http.StatusOK, newMemberID)
}

func handlePatchMember(w http.ResponseWriter, r *http.Request) {
	var params db.Member

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

	ok, err := updateMember(mux.Vars(r)["id"], params, image)

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrMemberNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}

func handleDeleteMember(w http.ResponseWriter, r *http.Request) {
	ok, err := deleteMember(mux.Vars(r)["id"])

	if err != nil {
		http_wrapper.Failure(w, http.StatusInternalServerError, err)
		return
	}

	if !ok {
		http_wrapper.Failure(w, http.StatusNotFound, ErrMemberNotFound)
		return
	}

	http_wrapper.Success(w, http.StatusNoContent, nil)
}
