package main

import (
	"fmt"
	"log"
	"net/http"
)

func abc(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 page not", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err : %v", err)
			return
		}

		fmt.Fprintf(w, "POST from website r.postform = %v\n", r.PostForm)
		name := r.FormValue("name")
		addr := r.FormValue("address")

		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", addr)

	default:
		fmt.Fprintf(w, "only get and post")
	}
}

func main() {

	addr := ":8080"
	http.HandleFunc("/", abc)
	log.Println("server has started on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("hello world")

	fmt.Fprintf(w, "Hello")
}
