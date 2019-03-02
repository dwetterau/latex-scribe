package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/dwetterau/latex-scribe/recognize"
)

func main() {
	rec := recognize.New()

	certFile := os.Getenv("cert_file")
	keyFile := os.Getenv("key_file")

	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static"))))
	http.HandleFunc("/submit-canvas-input", submitHandler(rec))
	log.Fatal(http.ListenAndServeTLS(":8080", certFile, keyFile, nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("../templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type submitBody struct {
	Data string `json:"data"`
}

func submitHandler(recognizer recognize.Recognizer) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
			return
		}
		b, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data submitBody
		err = json.Unmarshal(b, &data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		latex, err := recognizer.ToLatex(data.Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println("Recognized: ", latex)
		w.WriteHeader(200)
		w.Write([]byte(latex))
	}
}
