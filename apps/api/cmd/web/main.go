package main

import (
	"io"
	"log"
	"log/slog"
	"net/http"

	"github.com/j4y_funabashi/hikes/apps/api/pkg/app"
)

const (
	FitFileArchiveDir = "/archive"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func handleGpxUpload(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	logger := slog.Default()

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		logger.Error("failed parsing multipart form", "err", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	file, fileHeader, err := r.FormFile("file")
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

	app.ProcessFileUpload(fileBytes, fileHeader.Filename, FitFileArchiveDir)

	logger.Info("file uploaded", "name", fileHeader.Filename)
}

func main() {
	port := ":8080"
	slog.Info("Starting hikes server", "port", port)

	http.HandleFunc("/gpx", handleGpxUpload)
	log.Fatal(http.ListenAndServe(port, nil))
}
