package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func handleGpxUpload(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default()
	app := App{
		logger: logger,
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		logger.Error("failed parsing multipart form", "err", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		logger.Error("failed parsing form file", "err", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		logger.Error("failed reading file", "err", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	gpxFile, err := app.ProcessGPXFileUpload(fileBytes)
	if err != nil {
		logger.Error("failed processing gpx file upload", "err", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	logger.Info("file uploaded", "filename", gpxFile.Name())
}

type App struct {
	logger *slog.Logger
}

func (app App) ProcessGPXFileUpload(fileBytes []byte) (*os.File, error) {
	return saveFile(fileBytes)
}

func saveFile(fileBytes []byte) (*os.File, error) {
	tempFile, err := os.CreateTemp(".", "upload-*.gpx")
	if err != nil {
		return nil, fmt.Errorf("failed creating temp file: %s", err)
	}
	defer tempFile.Close()
	tempFile.Name()

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return nil, fmt.Errorf("failed writing to temp file: %s", err)
	}

	return tempFile, nil
}

func main() {
	port := ":8080"
	slog.Info("Starting hikes server", "port", port)

	http.HandleFunc("/gpx", handleGpxUpload)
	log.Fatal(http.ListenAndServe(port, nil))
}
