// package main

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"

// 	prerecorded "github.com/deepgram/deepgram-go-sdk/pkg/api/prerecorded/v1"
// 	interfaces "github.com/deepgram/deepgram-go-sdk/pkg/client/interfaces"
// 	client "github.com/deepgram/deepgram-go-sdk/pkg/client/prerecorded"
// )

// const googleKey = "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"
// const url = "https://api.openai.com/v1/chat/completions" // URL для чатовой модели

// const (
// 	apiKey   string = "69b15904923e6faba6b78f06aa0051d623ebb8e2"
// 	filePath string = "kz.mp3"
// )

// func main() {

// 	client.InitWithDefault()

// 	ctx := context.Background()

// 	options := interfaces.PreRecordedTranscriptionOptions{
// 		Model:       "nova-2",
// 		SmartFormat: true,
// 		Language:    "ru-RU", // Установка русского языка
// 	}

// 	// Получение ключа API из переменной среды
// 	apiKey := "69b15904923e6faba6b78f06aa0051d623ebb8e2"
// 	if apiKey == "" {
// 		fmt.Println("DEEPGRAM_API_KEY is not set")
// 		os.Exit(1)
// 	}

// 	// Инициализация клиента с ключом API
// 	c := client.New(apiKey, interfaces.ClientOptions{})
// 	dg := prerecorded.New(c)

// 	res, err := dg.FromFile(ctx, filePath, options)
// 	if err != nil {
// 		fmt.Printf("Error from Deepgram API: %v\n", err)
// 		os.Exit(1)
// 	}

// 	if res.Metadata.RequestID == "" {
// 		fmt.Println("No request ID received from Deepgram API")
// 		os.Exit(1)
// 	}

// 	// Извлечение текста из результатов
// 	var transcripts []string
// 	for _, channel := range res.Results.Channels {
// 		for _, alternative := range channel.Alternatives {
// 			if alternative.Transcript != "" {
// 				transcripts = append(transcripts, alternative.Transcript)
// 			}
// 		}
// 	}

// 	// Вывод текста на экран
// 	if len(transcripts) == 0 {
// 		fmt.Println("No transcripts received from Deepgram API")
// 		os.Exit(1)
// 	}

// 	var text string
// 	for _, transcript := range transcripts {
// 		fmt.Printf("Transcript: %s\n", transcript)
// 		text = transcript
// 	}

// 	// GPT

// 	data := map[string]interface{}{
// 		"messages": []map[string]interface{}{
// 			{"role": "user", "content": "Ответь на казахском языке" + text},
// 		},
// 		"model": "gpt-4-turbo",
// 	}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {

// 		return
// 	}

// 	request, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
// 	if err != nil {

// 		return
// 	}
// 	request.Header.Add("Content-Type", "application/json")
// 	request.Header.Add("Authorization", "Bearer "+googleKey)

// 	clientt := &http.Client{}
// 	response, err := clientt.Do(request)
// 	if err != nil {

// 		return
// 	}
// 	defer response.Body.Close()
// 	gptBody, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		return
// 	}

// 	var gptDataResponse map[string]interface{}
// 	if err := json.Unmarshal(gptBody, &gptDataResponse); err != nil {
// 		fmt.Println("Error decoding GPT response body:", err)
// 		os.Exit(1)
// 	}

// 	gptMessage := gptDataResponse["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)
// 	// VOICE

// 	apinKey := "gD8gnoW3hP8MqcGmf0f6c8K4WMk9QYOW7dDOZ6Qw"
// 	if apinKey == "" {
// 		log.Fatal("NARAKEET_API_KEY environment variable is not set")
// 	}

// 	voice := "Altynai"

// 	url := "https://api.narakeet.com/text-to-speech/mp3?voice=" + voice
// 	outputFilePath := "output.mp3"

// 	// Создание HTTP клиента и настройка запроса
// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(string(gptMessage))))
// 	if err != nil {
// 		log.Fatal("Error creating request: ", err)
// 	}

// 	// Установка заголовков
// 	req.Header.Set("Accept", "application/octet-stream")
// 	req.Header.Set("Content-Type", "text/plain")
// 	req.Header.Set("x-api-key", apinKey)

// 	// Выполнение запроса
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal("Error sending request: ", err)
// 	}
// 	defer resp.Body.Close()

// 	// Сохранение ответа в файл
// 	outputFile, err := os.Create(outputFilePath)
// 	if err != nil {
// 		log.Fatal("Error creating file: ", err)
// 	}
// 	defer outputFile.Close()

// 	_, err = io.Copy(outputFile, resp.Body)
// 	if err != nil {
// 		log.Fatal("Error writing to file: ", err)
// 	}

// }
