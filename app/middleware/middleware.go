package middleware

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/syrlramadhan/go-market/app/util"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		var token string

		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			cookie, err := r.Cookie("token")
			if err != nil {
				if r.URL.Path == "/assets/" {
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}else {
					http.Redirect(w, r, "/login", http.StatusSeeOther)
					return
				}
			}
			token = cookie.Value
		}
		isValid, err := util.ValidateToken(token)
		if err != nil || !isValid {
			if r.URL.Path == "/assets/" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}

		next(w, r, p)
	}
}