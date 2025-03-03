package main

import (
	"fmt"
	"net/http"

	"github.com/domainr/whois"
)

func main() {
	http.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
		domain := r.URL.Query().Get("domain")
		if domain == "" {
			http.Error(w, "Missing 'domain' parameter", http.StatusBadRequest)
			return
		}

		request, err := whois.NewRequest(domain)
		if err != nil {
			http.Error(w, "Invalid domain", http.StatusBadRequest)
			return
		}

		response, err := whois.DefaultClient.Fetch(request)
		if err != nil {
			http.Error(w, "WHOIS lookup failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s", response.Body)
	})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
