package auth

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"

	"github.com/titpetric/platform/registry"
)

func init() {
	// Register the plugin
	registry.Add("auth.RegisterAuth", RegisterAuth, nil)
}

func RegisterAuth(r chi.Router) {
	// Configure Goth providers
	goth.UseProviders(
		github.New("GITHUB_KEY", "GITHUB_SECRET", "http://localhost:8080/auth/github/callback"),
	)

	// Group: /auth
	r.Route("/auth", func(r chi.Router) {
		r.Get("/{provider}", func(w http.ResponseWriter, r *http.Request) {
			gothic.BeginAuthHandler(w, r)
		})

		r.Get("/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
			user, err := gothic.CompleteUserAuth(w, r)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			fmt.Fprintf(w, "User: %+v", user)
		})

		r.Get("/logout/{provider}", func(w http.ResponseWriter, r *http.Request) {
			gothic.Logout(w, r)
			fmt.Fprintln(w, "Logged out")
		})
	})
}
