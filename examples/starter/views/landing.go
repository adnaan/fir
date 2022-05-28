package views

import (
	"net/http"

	"github.com/adnaan/authn"
	pwc "github.com/adnaan/pineview/controller"
)

type LandingView struct {
	pwc.DefaultView
	Auth *authn.API
}

func (l *LandingView) Content() string {
	return "./templates/views/landing"
}

func (l *LandingView) Layout() string {
	return "./templates/layouts/index.html"
}

func (l *LandingView) OnMount(_ http.ResponseWriter, r *http.Request) (pwc.Status, pwc.M) {
	if r.Method != "GET" {
		return pwc.Status{Code: 405}, nil
	}
	if _, err := l.Auth.CurrentAccount(r); err != nil {
		return pwc.Status{Code: 200}, nil
	}
	return pwc.Status{Code: 200}, pwc.M{
		"is_logged_in": true,
	}
}
