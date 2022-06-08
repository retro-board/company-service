package company

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
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

	key := r.Header.Get("X-Auth-Key")
	if key == "" {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "X-Auth-Key is required"})
		return
	}
	userID := r.Header.Get("X-Auth-User")
	if userID == "" {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "X-Auth-User is required"})
		return
	}

	if err := c.verifyKey(key, userID); err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	type respValue struct {
		Exists    bool   `json:"exists"`
		Domain    string `json:"domain"`
		SubDomain string `json:"subdomain"`
		Name      string `json:"name"`
	}

	domain := chi.URLParam(r, "domain")
	if domain == "" {
		http.Error(w, "domain is required", http.StatusBadRequest)
		return
	}

	err := c.companyParts(domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exists, err = c.CheckCompanyExists(domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exists {
		jsonResponse(w, http.StatusOK, &respValue{
			Exists:    true,
			Domain:    c.CompanyAccount.Domain,
			SubDomain: c.CompanyAccount.Subdomain,
			Name:      c.CompanyAccount.Name,
		})
	} else {
		jsonResponse(w, http.StatusNotFound, &respValue{
			Exists:    false,
			Domain:    c.CompanyAccount.Domain,
			SubDomain: c.CompanyAccount.Subdomain,
			Name:      c.CompanyAccount.Name,
		})
	}
}
