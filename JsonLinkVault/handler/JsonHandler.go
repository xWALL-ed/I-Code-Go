package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type jsonDecode struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

func JsonHandler(filePath string, fallback http.Handler) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request){
		filePath := "JsonLinkVault\\url.json"
			jsonData, err := jsonToMap(filePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			path := r.URL.Path
			if dest, ok := jsonData[path]; ok {
				http.Redirect(w, r, dest, http.StatusFound)
				return 
			}
			fallback.ServeHTTP(w, r)
	}
}

func jsonToMap(filePath string) (map[string]string, error) {
	jsonURLs := make(map[string]string)
	
	jsonStore, err := parseJsonFile(filePath)
	if err != nil {
		return nil, err
	}

	for _, item := range jsonStore {
		jsonURLs[item.Path] = item.Url
	}

	return jsonURLs, nil
}

func parseJsonFile(filePath string) ([]jsonDecode, error) {
	jsonStore := []jsonDecode{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ERROR readJsonFile: Opening File: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if _ , err = decoder.Token(); err != nil {
		return nil, fmt.Errorf("ERROR readJsonFile: Reading 1st token: %v", err)
	}
	for decoder.More() {
		var item jsonDecode
		err := decoder.Decode(&item)
		if err != nil {
			return nil, fmt.Errorf("ERROR readJsonFile: Decoding: %v", err)
		}
		jsonStore = append(jsonStore, item)
	}
	if _ , err = decoder.Token(); err != nil {
		return nil, fmt.Errorf("ERROR readJsonFile: Reading last token: %v", err)
	}

	return jsonStore, nil
}