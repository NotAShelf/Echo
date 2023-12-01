package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
)

var version string

type PageData struct {
	Files   []string
	Version string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		basePath = "."
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if regexp.MustCompile(`\.\.\/?`).MatchString(r.URL.Path) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if r.URL.Path == "/" {
			files, err := os.ReadDir(basePath)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fileNames := make([]string, len(files))
			for i, file := range files {
				fileNames[i] = file.Name()
			}

			data := PageData{
				Files:   fileNames,
				Version: version,
			}

			tmpl := template.Must(template.ParseFiles("public/template.html"))
			tmpl.Execute(w, data)
		} else {
			filePath := filepath.Join(basePath, r.URL.Path)
			http.ServeFile(w, r, filePath)
		}
	})

	port, _ := strconv.Atoi(serverPort)
	fmt.Printf("Echo started with base path %s on port %d\n", basePath, port)
	http.ListenAndServe(":"+serverPort, nil)
}
