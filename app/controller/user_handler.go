package controller

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/syrlramadhan/go-market/app/model"
	"github.com/syrlramadhan/go-market/app/util"
	"gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var id = uuid.New()
		var firstName = r.FormValue("first-name")
		var lastName = r.FormValue("last-name")
		var email = r.FormValue("email")
		var password = r.FormValue("password")

		if firstName == "" || email == "" || password == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		password, err := util.HashPassword(password)
		if err != nil {
			http.Error(w, "error while password is hashed", http.StatusInternalServerError)
		}

		user := model.MstUser{
			Id:        id.String(),
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  password,
		}

		if err := CreateUser(user, db); err != nil {
			http.Error(w, "Error while creating user", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func LoginUserHandler(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// userId := p.ByName("id")

		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := GetUserByEmail(email, db)
		emailError := fmt.Sprintf("user with email %s not found", email)
		if err != nil {
			http.Error(w, emailError, http.StatusInternalServerError)
			return
		}

		if util.ComparePassword(user.Password, password) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			http.Error(w, "incorrect password", http.StatusInternalServerError)
		}
	}
}