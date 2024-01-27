package auth

import (
    "crypto/sha256"
    "crypto/subtle"
    "net/http"
	"super.com/networking/data"
)

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			
			value, exists := data.RegisteredUsers[username]

			if !exists {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				http.Error(w, "Create user", http.StatusUnauthorized)
				return
			}

			expectedPasswordHash := sha256.Sum256([]byte(value))

			passwordHash := sha256.Sum256([]byte(password))

			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if exists && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})

} 