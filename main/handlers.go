package main

import (
	"html/template"
	"log"
	"net/http"
)

// Обработчик главной странице.
func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"Kurs_sem_5\\Kurs_sem5\\ui\\html\\home.page.tmpl.xhtml",
		//"Kurs_sem5/ui/html/base.layout.tmpl.xhtml",
		//"Kurs_sem5/ui/html/footer.partial.tmpl.xhtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
