package cookie

import (
	"net/http"
	"time"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request, token string) {
	cookie := http.Cookie{
		Name:    "remember_me",
		Value:   token,
		Expires: time.Now().Add(7 * 24 * time.Hour),
		Path:    "/",
	}

	http.SetCookie(w, &cookie)
	w.Write([]byte("Remember cookie telah diatur"))
}
