package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetDownloadFolders(downloadFolders map[string]string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var folders []string
		for folder := range downloadFolders {
			folders = append(folders, folder)
		}
		response, err := json.Marshal(folders)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, string(response))
	}
}
