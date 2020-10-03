package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody return the status
type HealthBody struct {
	Status string `json:"status"`
}

// Health GET entrypoint
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var status = HealthBody{"alive"}

	_ = json.NewEncoder(w).Encode(status)
}
