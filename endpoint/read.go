package endpoint

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the cocktail name from the query parameters.
	cocktail := r.URL.Query().Get("cocktail")
	if cocktail == "" {
		http.Error(w, "Missing cocktail parameter", http.StatusBadRequest)
		return
	}
	// Get the base API URL from .env.
	baseURL := os.Getenv("CocktailSearchURL")
	if baseURL == "" {
		http.Error(w, "CocktailSearchURL not set in environment", http.StatusInternalServerError)
		return
	}
	// Build the external URL using the provided cocktail name.
	apiURL := baseURL + url.QueryEscape(cocktail)

	resp, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error fetching data from API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}
