package api

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

type Cocktail struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Info string `json:"info"`
}

// APIHandler forwards requests to the external CocktailDB API.
func APIHandler(w http.ResponseWriter, r *http.Request) {
	queryS := r.URL.Query().Get("s")
	// Read base API URL from .env as CocktailSearchURL.
	baseURL := os.Getenv("CocktailSearchURL")
	if baseURL == "" {
		http.Error(w, "CocktailSearchURL not set in environment", http.StatusInternalServerError)
		return
	}
	// Build the external URL using the base URL and query parameter.
	apiURL := baseURL + url.QueryEscape(queryS)

	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}
