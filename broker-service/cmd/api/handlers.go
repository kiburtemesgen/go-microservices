package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {

	payload := jsonRespones{
		Error:   false,
		Message: "hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}
