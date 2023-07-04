# ElevenLabs Voice Cloning API

Subscription to the ElevenLabs Voice Cloning application allows you to use the API to create Audio speech from any given text.

I have built this application to make the API calls in Golang. You will find similiar examples in Python within the Elevenlabs.io site.

Steps needed

1) Get an Elevenlabs API Key which is found any your profile on the site
   
2) Export the API Key as an environment variable:
   `export ELEVENLABS_API_KEY=<your API key>`
   
3) Check the config.yaml file and change any settings in there. You will need a Voice ID to change which voice you want to use. The text that you wanted converted to audio is also there within the config.yaml file

To run the application and convert the text to an audio speech MP3 file:

```go run main.go```
