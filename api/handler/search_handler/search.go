package search_handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func PostSearch(w http.ResponseWriter, r *http.Request) {
	body := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&body)
	w.Header().Set("content-type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		return
	}
	query := `{
		"search_type": "match",
		"query": {
			"term": "`
	query = query + body["param"]
	query = query + `",
			"start_time": "2000-01-01T00:00:00.000Z",
			"end_time": "2030-01-01T00:00:00.000Z"
		},
		"from": 0,
		"max_results": 20,
		"_source": []
	}`
	req, err := http.NewRequest("POST", "http://localhost:4080/api/default/_search", strings.NewReader(query))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		return
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		return
	}
	defer resp.Body.Close()
	bod, err := io.ReadAll(resp.Body) 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		return
	}
	response := map[string]interface{}{}
	err = json.Unmarshal(bod, &response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Internal server error"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
