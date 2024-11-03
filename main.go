package main

import (
	"fmt"
	"html/template"
	"time"

	"log"
	"net/http"
)

type Film struct {
	Title    string
	Director string
}

//learning
//creating go webserver / htmx integration / template

func main() {
	fmt.Println("hello world")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, "hello world")
		// io.WriteString(w, r.Method)
		time.Sleep(3 * time.Second)
		tmp1 := template.Must(template.ParseFiles("index.html")) //rendering
		films := map[string][]Film{
			"Films": {
				{Title: "film1", Director: "franics ford coppola"},
				{Title: "film2", Director: "franics ford coppola"},
				{Title: "film3", Director: "franics ford coppola"},
			},
		}
		tmp1.Execute(w, films)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		// log.Print("HTMX request received")
		// log.Print(r.Header.Get("HX_Request"))
		title := r.PostFormValue("title")
		director := r.PostFormValue("director")
		htmlstr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
		tmpl, _ := template.New("t").Parse(htmlstr)
		tmpl.Execute(w, nil)
		//use template fragmention task
	}
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-film/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
