package helpers

import (
	"crypto/rand"
	"io"
	"os"
	"path"
	"strings"
)

func CreateLocalSourcePath(workDir string, src string, ID string) string {
	// Get local destination path.
	tmpDir := workDir + "/" + ID + "/"
	os.MkdirAll(tmpDir, 0700)
	os.MkdirAll(tmpDir+"src", 0700)
	os.MkdirAll(tmpDir+"dst", 0700)
	slices := strings.Split(src, ".")
	ext := slices[len(slices)-1]
	return tmpDir + path.Base(ID+"."+ext)
}

func GetTmpPath(workDir string, ID string) string {
	tmpDir := workDir + "/" + ID + "/"
	return tmpDir
}

func GenerateRandomKey(length int) []byte {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}
