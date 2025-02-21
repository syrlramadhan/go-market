package util

import "net/http"

func IsLogin(r *http.Request) bool {
	cookie, err := r.Cookie("token")
	isAuthenticated := false

	if err == nil && cookie != nil {
		// Validasi token
		isValid, _ := ValidateToken(cookie.Value)
		isAuthenticated = isValid
	}

	return isAuthenticated
}
