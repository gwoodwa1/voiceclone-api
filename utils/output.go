package utils

import (
	"os"
	"net/http"
	"io"
	"fmt"

)

const (
	ChunkSize = 1024
)

func WriteToOutputFile(response *http.Response, outputFilePath string) error {
	outFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("error creating output file: %w", err)
	}
	defer outFile.Close()

	buffer := make([]byte, ChunkSize)
	for {
		n, err := response.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("error reading response: %w", err)
		}

		if _, err := outFile.Write(buffer[:n]); err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}
	return nil
}
