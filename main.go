package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"github.com/gwoodwa1/voiceclone-api/utils"
)




func main() {
    
	config := utils.LoadConfig("./config.yaml")

	url := fmt.Sprintf("https://api.elevenlabs.io/v1/text-to-speech/%s", config.Settings.VoiceID)
	
	headers := map[string]string{
		"Accept":       "audio/mpeg",
		"Content-Type": "application/json",
		"xi-api-key":   config.ApiKey,
	}

	data := map[string]interface{}{
		"text":     config.TextToAudio,
		"model_id": "eleven_monolingual_v1",
		"voice_settings": map[string]float64{
			"stability":       config.Settings.Stability,
			"similarity_boost": config.Settings.SimilarityBoost,
		},
	}

	dataBytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error encoding data to JSON: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataBytes))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for headerKey, headerValue := range headers {
		req.Header.Add(headerKey, headerValue)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()
    
	if err := utils.WriteToOutputFile(resp, config.OutputFilePath); err != nil{
		log.Fatal("output file not created properly", err)
	}

}
