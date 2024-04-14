package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func synonymHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.FormValue("text")

	data := map[string]interface{}{
		"messages": []map[string]interface{}{
			{"role": "user", "content": "Найди синонимы на казахском языке и напиши их: " + text},
		},
		"model": "gpt-4-turbo",
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

	var response CompletionResponse
	if err := json.Unmarshal(body, &response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(response.Choices) > 0 {
		fmt.Fprintf(w, response.Choices[0].Message.Content)
	} else {
		fmt.Fprint(w, "Ошибка вывода")
	}
}
