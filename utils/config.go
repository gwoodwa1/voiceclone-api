package utils

import(
	"log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	
)

type Config struct {
	TextToAudio    string `yaml:"textToAudio"`
	ApiKey         string `yaml:"apiKey"`
	OutputFilePath string `yaml:"outputFilePath"`
	Settings       struct {
		Stability       float64 `yaml:"stability"`
		SimilarityBoost float64 `yaml:"similarityBoost"`
		VoiceID         string  `yaml:"voiceID"`
	} `yaml:"settings"`
}

func LoadConfig(filePath string) Config {	// Load config
	var config Config
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}
	return config
}
