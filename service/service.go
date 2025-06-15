package service

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetMethod(url string) (map[string]any, error) {
	var res, err = http.Get(url)
	defer res.Body.Close()

	if err != nil {
		return map[string]any{}, err
	}

	by, err := io.ReadAll(res.Body)
	if err != nil {
		return map[string]any{}, err
	}

	var data map[string]any
	err = json.Unmarshal(by, &data)

	if err != nil {
		return map[string]any{}, err
	}

	return data, nil
}
