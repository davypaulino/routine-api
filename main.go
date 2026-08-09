package main

import	(
	"fmt"
	"time"
	"net/http"
	"os"
	"path/filepath"

	"github.com/davypaulino/routine-api/domain"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		response := fmt.Sprintf(`{"status":"ok","time":"%s"}`, domain.NewGoal("test", "test", time.Now()))
		w.Write([]byte(response))
	})

	mux.HandleFunc("GET /api-docs", func(w http.ResponseWriter, r *http.Request) {
		pwd, err := os.Getwd()
		if err != nil {
			http.Error(w, "Failed to get current working directory", http.StatusInternalServerError)
			return
		}

		apiDocsPath := filepath.Join(pwd, "api-docs/index.html")
		response, err := os.ReadFile(apiDocsPath)
		if err == nil && response != nil {
			w.WriteHeader(http.StatusOK)
			w.Write(response)
			return
		}

		http.Error(w, "API documentation not found", http.StatusNotFound)
		return
	})

	http.ListenAndServe(":8080", mux)
}
