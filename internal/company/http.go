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

	type respValue struct {
		Exists    bool   `json:"exists"`
		Domain    string `json:"domain"`
		SubDomain string `json:"subdomain"`
		Name      string `json:"name"`
	}

	domain := r.Header.Get("domain")
	if domain == "" {
		http.Error(w, "domain header is required", http.StatusBadRequest)
		return
	}

	err := c.companyParts(domain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exists, err = c.checkCompanyExists(domain)
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
	return
}
