package company

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c Company) CreateHandler(w http.ResponseWriter, r *http.Request) {

}

func (c Company) ViewHandler(w http.ResponseWriter, r *http.Request) {

}

func (c Company) UpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func (c Company) ExistsHandler(w http.ResponseWriter, r *http.Request) {
	var exists bool

	companyName := r.Header.Get("Company-Name")
	if companyName == "" {
		http.Error(w, "Company-Name header is required", http.StatusBadRequest)
		return
	}

	exists, err := c.checkCompanyExists(companyName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		jsonResponse(w, http.StatusOK, struct {
			Exists bool `json:"exists"`
		}{
			Exists: exists,
		})
	} else {
		jsonResponse(w, http.StatusNotFound, struct {
			Exists bool `json:"exists"`
		}{
			Exists: exists,
		})
	}
	return
}
