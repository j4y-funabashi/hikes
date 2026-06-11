package app_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/j4y_funabashi/hikes/apps/api/pkg/app"
	"gotest.tools/v3/assert"
)

func TestProcessFileUpload(t *testing.T) {
	testFilepath := "testdata/test_1.fit"
	fileData, err := os.ReadFile(testFilepath)
	assert.NilError(t, err)

	h := md5.Sum(fileData)
	hash := hex.EncodeToString(h[:])
	t.Log(hash)

	archiveDir := "/archive"
	err = app.ProcessFileUpload(fileData, "test_1.fit", archiveDir)
	assert.NilError(t, err)

	expectedArchivedFitFilepath := filepath.Join(archiveDir, "fit", fmt.Sprintf("%s.fit", hash))
	_, err = os.Stat(expectedArchivedFitFilepath)
	assert.NilError(t, err)

	expectedArchivedGpxFilepath := filepath.Join(archiveDir, "gpx", fmt.Sprintf("%s.gpx", hash))
	_, err = os.Stat(expectedArchivedGpxFilepath)
	assert.NilError(t, err)
}
