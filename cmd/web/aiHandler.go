package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "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"
const url = "https://api.openai.com/v1/chat/completions" // URL для чатовой модели

func CorrectHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.FormValue("text")

	data := map[string]interface{}{
		"messages": []map[string]interface{}{
			{"role": "user", "content": "Скорректируйте текст, чтобы он был грамматически правильным на казахском языке: " + text},
		},
		"model": "gpt-3.5-turbo",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<h1>Исходный текст</h1><p>%s</p><h1>Скорректированный текст</h1><p>%s</p>", text, string(body))
}
