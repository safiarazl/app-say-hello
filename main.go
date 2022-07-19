package main

import(
	"fmt"
	go_say_hello "github.com/safiarazl/go-say-hello"
	"time"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// fmt.Fprintf(w, "Form Data: %v", r.Form)
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")
	fmt.Fprintf(w, "First Name: %s\n", fname)
	fmt.Fprintf(w, "Last Name: %s\n", lname)
}

func main()  {
	fileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}


	fmt.Println(go_say_hello.SayHello())
	fmt.Println(time.Now())
	fmt.Println(time.Since(time.Date(2022, time.July, 18, 0, 0, 0, 0, time.UTC)))
}