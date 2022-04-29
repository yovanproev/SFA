package hns

import (
	"hns/hns"
	"log"
	"net/http"
	"text/template"
)

func IndexTemplate(gotFirstTen hns.SummedType) {
	mux := http.NewServeMux()

	templates := populateTemplates()

	mux.HandleFunc("/top", func(w http.ResponseWriter, r *http.Request) {
		t := templates.Lookup("index.html")
		t.Execute(w, gotFirstTen)

		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	})
	http.ListenAndServe(":9000", mux)
}

func populateTemplates() *template.Template {
	result := template.New("")
	const basePath = "Week10/Lecture27/templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))

	return result
}
