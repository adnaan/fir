package app

import (
	"net/http"

	"github.com/adnaan/authn"
	"github.com/adnaan/fir"
)

type DashboardView struct {
	fir.DefaultView
	Auth *authn.API
}

func (d *DashboardView) Content() string {
	return "./templates/views/app"
}

func (d *DashboardView) Layout() string {
	return "./templates/layouts/app.html"
}

func (d *DashboardView) OnGet(w http.ResponseWriter, r *http.Request) fir.Page {
	return fir.Page{
		Data: fir.Data{
			"is_logged_in": true,
		}}
}
