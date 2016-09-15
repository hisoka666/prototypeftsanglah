package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"appengine"
	"appengine/user"
)

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/main", mainPage)
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/searchCM", searchCM)
	http.HandleFunc("/tambahdata", tambahdata)
	http.HandleFunc("/welcome", welcome)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(ctx, "/")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	s1, _ := template.ParseFiles("templates/base.html", "templates/index.html")
	s1.Execute(w, nil)
}

func searchCM(w http.ResponseWriter, r *http.Request) {
	nocm := r.PostFormValue("nocm")
	fmt.Fprint(w, nocm)
}

//TODO:fungsi login
func login(w http.ResponseWriter, r *http.Request) {

}

//TODO:fungsi tambah data pasien

func tambahdata(w http.ResponseWriter, r *http.Request) {

}

//TODO:fungsi registrasi
func register(w http.ResponseWriter, r *http.Request) {

}

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/main" {
		http.NotFound(w, r)
		return
	}
	s1, _ := template.ParseFiles("templates/base.html", "templates/main.html")
	s1.Execute(w, nil)
}
