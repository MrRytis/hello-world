package handler

import (
	"github.com/MrRytis/hello-world/utils"
	"html/template"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type PageData struct {
	Name string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	hello := Response{
		Message: "Hello World!",
	}

	utils.JSON(w, http.StatusOK, hello)
}

func Index(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "World"
	}

	data := PageData{
		Name: name,
	}

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to text/html
	w.Header().Set("Content-Type", "text/html")

	// Render the HTML template with the provided data and write it to the response
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
