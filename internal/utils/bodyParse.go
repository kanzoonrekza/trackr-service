package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"trackr-service/internal/types"
)

func BodyParseJson(r *http.Request) (types.JSON, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return types.JSON{}, err
	}
	defer r.Body.Close()

	var data types.JSON
	if err := json.Unmarshal(body, &data); err != nil {
		return types.JSON{}, err
	}

	return data, nil
}

func BodyParseFormData(r *http.Request) (types.FormDataFields, types.FormDataFile, error) {
	// 32 MB is the default used by ParseMultipartForm
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return types.FormDataFields{}, types.FormDataFile{}, err
	}

	formData := make(types.FormDataFields)

	for key, values := range r.MultipartForm.Value {
		formData[key] = values[0] // Assuming you only need the first value for each key
	}

	formDataFile := make(types.FormDataFile)

	for key, fileHeaders := range r.MultipartForm.File {
		formDataFile[key] = fileHeaders
	}

	return formData, formDataFile, nil
}
