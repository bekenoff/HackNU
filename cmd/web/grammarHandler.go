package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "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"
const url = "https://api.openai.com/v1/chat/completions"

type CompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
}

func CorrectHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.FormValue("text")

	data := map[string]interface{}{
		"messages": []map[string]interface{}{
			{"role": "user", "content": "Correct the text so that it is grammatically correct in the Kazakh language and indicate errors that were made in the grammar and give answer only in Kazakh language: " + text},
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
