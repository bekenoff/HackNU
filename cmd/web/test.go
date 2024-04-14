package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Question struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	Answer   string `json:"answer,omitempty"`
	Category string `json:"category"`
}

var db *sql.DB

func getAllQuestions() ([]Question, error) {
	rows, err := db.Query("SELECT id, question, category FROM ainur_hacknu.questions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []Question

	for rows.Next() {
		var question Question
		if err := rows.Scan(&question.ID, &question.Text, &question.Category); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}

func handleAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := getAllQuestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(questions)
}

func handleAnswer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var answer struct {
		ID     string `json:"id"`
		Answer string `json:"answer"`
	}

	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var correctAnswer string
	err = db.QueryRow("SELECT answer FROM ainur_hacknu.questions WHERE id = ?", answer.ID).Scan(&correctAnswer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Correct bool   `json:"correct"`
		Message string `json:"message"`
	}{
		Correct: answer.Answer == correctAnswer,
	}

	if response.Correct {
		response.Message = "Правильный ответ!"
	} else {
		response.Message = "Неправильный ответ!"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
