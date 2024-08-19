package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"trackr-service/internal/types"
)

func BodyParseJson(r *http.Request) (types.JSON, error) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		return types.JSON{}, errors.New("request Content-Type isn't application/json")
	}

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

func BodyParseFormData(r *http.Request) (types.JSON, types.FormDataFile, error) {
	// 32 MB is the default used by ParseMultipartForm
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return types.JSON{}, types.FormDataFile{}, err
	}

	formData := make(types.JSON)

	for key, values := range r.MultipartForm.Value {
		var parsedValue interface{}
		if err := json.Unmarshal([]byte(values[0]), &parsedValue); err != nil {
			// If parsing fails, treat it as a string
			formData[key] = values[0]
		} else {
			// If parsing succeeds, store the parsed JSON
			formData[key] = parsedValue
		}

	}

	formDataFile := make(types.FormDataFile)

	for key, fileHeaders := range r.MultipartForm.File {
		formDataFile[key] = fileHeaders
	}

	return formData, formDataFile, nil
}
