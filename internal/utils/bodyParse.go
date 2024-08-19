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

func BodyParseFormData(r *http.Request) (types.FormData, error) {
	// 32 MB is the default used by ParseMultipartForm
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return types.FormData{}, err
	}

	formData := make(types.FormData)

	for key, values := range r.MultipartForm.Value {
		formData[key] = values[0] // Assuming you only need the first value for each key
	}

	return formData, nil
}
