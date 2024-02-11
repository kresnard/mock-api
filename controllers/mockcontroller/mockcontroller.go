package mockcontroller

import (
	"mock_api/models/mockmodel"
	"net/http"
	"path"
	"strings"
)

func GetMock(w http.ResponseWriter, r *http.Request) {
	endpoint := path.Base(r.URL.Path)

	endpoint = strings.TrimSuffix(endpoint, "/")

	mockData, err := mockmodel.Get(endpoint)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	mockDataBytes := []byte(mockData)

	// Write the mock response to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write(mockDataBytes)
}
