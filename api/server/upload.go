package server

import (
	"log"
	"net/http"
)

type response struct {
	Message string     `json:"message"`
	Status  int        `json:"status"`
	FileName string `json:"file_name"`
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
  const MAX_UPLOAD_SIZE = 5 * 1024 * 1024 * 1024
  
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

  r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big.", http.StatusBadRequest)
		return
	}
  
  // The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
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
	dst, err := os.Create(fmt.Sprintf("/tmp/uploads/%d%s", outputFileName)))
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
	resp := response{
 		Message: "Uploaded",
		Status:  200,
		FileName: outputFileName,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode()
}
