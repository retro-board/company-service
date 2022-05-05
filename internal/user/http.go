package user

import (
	"fmt"
	"net/http"
	"time"
)

func (u User) callbackCookie(w http.ResponseWriter, r *http.Request, name, v string) {
	http.SetCookie(w, &http.Cookie{
		Name:   fmt.Sprintf("retro_%s", name),
		Value:  v,
		MaxAge: int(time.Hour.Seconds()),
		Secure: r.TLS != nil,
		Domain: r.Host,
	})
}

func (u User) CallbackHandler(w http.ResponseWriter, r *http.Request) {

}

func (u User) LoginHandler(w http.ResponseWriter, r *http.Request) {

}

func (u User) VerifyHandler(w http.ResponseWriter, r *http.Request) {

}

func (u User) PermissionHandler(w http.ResponseWriter, r *http.Request) {

}
