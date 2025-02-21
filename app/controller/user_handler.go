package controller

import (
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
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusInternalServerError)
			return
		}

		if !util.ComparePassword(user.Password, password) {
			http.Error(w, "Invalid email or password", http.StatusInternalServerError)
		}

		token, err := util.GenerateJWT(user.FirstName, user.LastName)
		if err != nil {
			http.Error(w, "failed to generate token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   600,
		})
		
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func LogoutUserHandler(db *gorm.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    "",
			Path:     "/",
			HttpOnly: true,
			MaxAge:   -1,
		})
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	
}
