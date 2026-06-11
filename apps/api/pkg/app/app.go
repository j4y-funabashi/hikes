package app

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// save fit file to archive dir
// convert to gpx + save to archive dir
// parse gpx file
// create static map
// save to db
func ProcessFileUpload(fileBytes []byte, fileName string, archiveDir string) error {

	fitFilename, err := saveFile(fileBytes, archiveDir)
	if err != nil {
		return err
	}

	err = convertFitFile(fitFilename, archiveDir)

	return err
}

func saveFile(fileBytes []byte, archiveDir string) (string, error) {
	fitArchiveDir := filepath.Join(archiveDir, "fit")
	err := os.MkdirAll(fitArchiveDir, 0777)
	if err != nil {
		return "", err
	}

	h := md5.Sum(fileBytes)
	hash := hex.EncodeToString(h[:])
	newFilename := fmt.Sprintf("%s.fit", hash)
	newFilepath := filepath.Join(fitArchiveDir, newFilename)
	newFile, err := os.Create(newFilepath)
	if err != nil {
		return "", err
	}

	_, err = newFile.Write(fileBytes)
	if err != nil {
		return "", err
	}

	return newFilepath, nil
}

func convertFitFile(fitFileName, archiveDir string) error {
	gpxArchiveDir := filepath.Join(archiveDir, "gpx")
	err := os.MkdirAll(gpxArchiveDir, 0777)
	if err != nil {
		return err
	}

	gpxFilename := filepath.Join(gpxArchiveDir, strings.TrimSuffix(filepath.Base(fitFileName), filepath.Ext(fitFileName))+".gpx")

	params := []string{
		"-i",
		"garmin_fit",
		"-f",
		fitFileName,
		"-o",
		"gpx",
		"-F",
		gpxFilename,
	}
	cmd := exec.Command("gpsbabel", params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
