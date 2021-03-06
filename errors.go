package fir

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func getUserMessage(status int, userMessage []string) string {
	msg := http.StatusText(status)
	if len(userMessage) > 0 {
		msg = strings.Join(userMessage, " ")
	}
	return msg
}

func ErrInternalServer(err error, userMessage ...string) Page {
	return Page{
		Code:    http.StatusInternalServerError,
		Message: getUserMessage(http.StatusInternalServerError, userMessage),
		Error:   err,
	}
}

func ErrBadRequest(err error, userMessage ...string) Page {
	return Page{
		Code:    http.StatusBadRequest,
		Message: getUserMessage(http.StatusBadRequest, userMessage),
		Error:   err,
	}
}

func ErrNotFound(err error, userMessage ...string) Page {
	return Page{
		Code:    http.StatusNotFound,
		Message: getUserMessage(http.StatusNotFound, userMessage),
		Error:   err,
	}
}

func morphError(err string) Patch {
	return Morph{
		Selector: "#fir-error",
		Template: &Template{
			Name: "fir-error",
			Data: Data{"error": err}},
	}
}

func PatchError(err error, userMessage ...string) Patchset {
	msg := "internal error"
	if err != nil && len(userMessage) == 0 {
		msg = err.Error()
		log.Printf("[controller] patch error: %s\n", err)
	}
	if len(userMessage) > 0 {
		msg = strings.Join(userMessage, " ")
		log.Printf("[controller] patch error: %s, message: %s\n", err, msg)
	}
	log.Printf("[controller] patch error: %s, %s\n", err, msg)
	return Patchset{morphError(msg)}
}

func PageError(err error, userMessage ...string) Page {
	msg := "internal error"
	if err != nil && len(userMessage) == 0 {
		msg = err.Error()
		log.Printf("[controller] page error: %s\n", err)
	}
	if len(userMessage) > 0 {
		msg = strings.Join(userMessage, " ")
		log.Printf("[controller] page error: %s, message: %s\n", err, msg)
	}

	data := Data{"error": msg}
	if msg == "" {
		data = nil
	}
	return Page{
		Code:  http.StatusOK,
		Error: err,
		Data:  data,
	}
}

func UnsetFormErrors(fields ...string) Patchset {
	var patchset Patchset

	for _, field := range fields {
		patchset = append(patchset, Morph{
			Selector: fmt.Sprintf("#%s-error", field),
			Template: &Template{
				Name: fmt.Sprintf("%s-error", field),
				Data: Data{
					fmt.Sprintf("#%sError", field): "",
				},
			},
		})
	}

	return patchset
}
