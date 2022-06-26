package todos

import (
	"net/http"

	"github.com/adnaan/fir"
	"github.com/adnaan/fir/cli/testdata/todos/models"
)

type View struct {
	DB *models.Client
	fir.DefaultView
}

func (v *View) Content() string {
	return "./views/todos"
}

func (v *View) Layout() string {
	return "./templates/layouts/index.html"
}

func (v *View) OnRequest(w http.ResponseWriter, r *http.Request) (fir.Status, fir.Data) {
	return fir.Status{Code: http.StatusOK}, nil
}