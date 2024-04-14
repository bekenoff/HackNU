package main

import (
	"HackNU/pkg/models"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func (app *application) fileUpload(w http.ResponseWriter, r *http.Request) {

	client_id := r.URL.Query().Get(":id")

	if client_id == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("video")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := "./uploads/" + header.Filename
	if _, err := os.Stat(filePath); err == nil {
		http.Error(w, "File already exists", http.StatusBadRequest)
		return
	}
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create the file for writing", http.StatusInternalServerError)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	newVideo := models.Video{
		Filename: header.Filename,
		Filepath: filePath,
		Likes:    0,
		ClientId: client_id,
	}

	err = app.video.Insert(&newVideo)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *application) showAllVideos(w http.ResponseWriter, r *http.Request) {
	videos, err := app.video.GetAllVideos()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(videos)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) incrementLike(w http.ResponseWriter, r *http.Request) {
	videoID := r.URL.Query().Get("video_id")
	clientID := r.URL.Query().Get("client_id")

	if videoID == "" || clientID == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err := app.video.IncrementLike(videoID, clientID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *application) getVideosWithLikes(w http.ResponseWriter, r *http.Request) {
	videos, err := app.video.GetAllVideosWithLikes()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(videos); err != nil {
		app.serverError(w, err)
		return
	}
}
