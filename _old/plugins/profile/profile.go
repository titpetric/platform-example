package profile

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"

	"github.com/titpetric/platform/registry"
)

func init() {
	registry.Add("profile.RegisterProfile", RegisterProfile, nil)
}

func RegisterProfile(r chi.Router) {
	r.Get("/profile", Handler())
}

func Handler() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			http.Error(w, "Not logged in", http.StatusUnauthorized)
			return
		}

		fmt.Fprintf(w, "Logged in user: %+v", user)
	}

	return http.HandlerFunc(handler)
}
