package accounts

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/adnaan/authn"
	"github.com/adnaan/fir"
)

type ConfirmView struct {
	fir.DefaultView
	Auth *authn.API
}

func (h *ConfirmView) Content() string {
	return "./templates/views/accounts/confirm"
}

func (h *ConfirmView) Layout() string {
	return "./templates/layouts/index.html"
}

func (h *ConfirmView) OnGet(w http.ResponseWriter, r *http.Request) fir.Page {
	token := chi.URLParam(r, "token")
	err := h.Auth.ConfirmSignupEmail(r.Context(), token)
	if err != nil {
		log.Println("err confirm.OnGet", err)
		return fir.Page{}
	}
	return fir.Page{
		Data: fir.Data{
			"confirmed": true,
		}}
}
