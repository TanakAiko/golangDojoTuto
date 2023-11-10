package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServeTLS("", "cert.pem", "key.pem", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

var userDB = map[string]string{
	"azba": "DZE",
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/login.html")
	if err != nil {
		fmt.Println("Error parsing", err)
		return
	}
	err = t.ExecuteTemplate(w, "login.html", "Please Log in")
	if err != nil {
		fmt.Println("Error executing template ", err)
		return
	}
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Goooooooooooood Joooooooooooob !")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Bad username or password!")
	}

}
