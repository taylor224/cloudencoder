package server

import (
	"net/http"
	"os"
	"io"
	"fmt"
	"time"
	"encoding/json"
	"path/filepath"
	
	"github.com/gin-gonic/gin"
)

type uploadResponse struct {
	Message string     `json:"message"`
	Status  int        `json:"status"`
	FileName string    `json:"file_name"`
}

func uploadHandler(c *gin.Context) {
    const MAX_UPLOAD_SIZE = 5 * 1024 * 1024 * 1024

	var w http.ResponseWriter = c.Writer
	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, MAX_UPLOAD_SIZE)

	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("/tmp/uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	outputFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename))
	dst, err := os.Create(fmt.Sprintf("/tmp/uploads/%s", outputFileName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Upload successful")
	resp := uploadResponse{
		Message: "Uploaded",
		Status:  200,
		FileName: outputFileName,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
