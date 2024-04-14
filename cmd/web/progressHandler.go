package main

import (
	"HackNU/pkg/models"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

func (app *application) createProgress(w http.ResponseWriter, r *http.Request) {
	var newProgress models.Progress

	body, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	err := json.NewDecoder(r.Body).Decode(&newProgress)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.progress.Insert(&newProgress)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) getProgress(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	progressData, err := app.progress.GetProgressById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.clientError(w, http.StatusNotFound)
		} else {
			app.serverError(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(progressData)
}

func (app *application) updateProgressLevel(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.progress.UpdateLevel(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) updateProgressPoints(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.progress.UpdatePoints(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) updateProgressTests(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.progress.UpdateTests(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) updateProgressFilms(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.progress.UpdateFilms(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) updateProgressMeetings(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")

	if id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.progress.UpdateMeetings(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteProgress(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.progress.DeleteProgressById(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
