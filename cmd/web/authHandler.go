package main

import (
	"HackNU/pkg/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (app *application) signupClient(w http.ResponseWriter, r *http.Request) {
	var newClient models.Client

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newClient)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.client.Insert(newClient.ClientMail, newClient.ClientPass)
	if err != nil {
		app.serverError(w, err)
		return
	}

	clientId, err := app.client.GetLastUserId()

	progressInstance := models.Progress{
		Level:    "A1",
		Points:   0,
		Tests:    0,
		Films:    0,
		Meetings: 0,
		ClientId: clientId,
	}

	err = app.progress.Insert(&progressInstance)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated) // 201
}

func (app *application) loginClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&client)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	clientId, err := app.client.Authenticate(client.ClientMail, client.ClientPass)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			app.clientError(w, http.StatusBadRequest)
			return
		} else {
			app.serverError(w, err)
			return
		}
	}

	responseUser, err := app.client.GetUserById(clientId)

	_, err = w.Write(responseUser)
	if err != nil {
		return
	}
}
