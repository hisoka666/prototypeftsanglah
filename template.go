package main

import (
   "log"
   "fmt"
   "net/http"
   "html/template"
)

func main(){
http.HandleFunc("/", index)
http.HandleFunc("/main", mainPage)
http.HandleFunc("/login", login)
http.HandleFunc("/register", register)
http.HandleFunc("/searchCM", searchCM)
http.HandleFunc("/tambahdata", tambahdata)
log.Print("Listening on port 8080")
log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    http.NotFound(w, r)
    return
  }
  s1, _ := template.ParseFiles("templates/base.html", "templates/index.html")
  s1.Execute(w, nil)
}

func searchCM(w http.ResponseWriter, r *http.Request){
  nocm := r.PostFormValue("nocm")
  fmt.Fprint(w, nocm)
}
//TODO:fungsi login
func login(w http.ResponseWriter, r *http.Request){


}

//TODO:fungsi tambah data pasien

func tambahdata(w http.ResponseWriter, r *http.Request){


}
//TODO:fungsi registrasi
func register(w http.ResponseWriter, r *http.Request){

}

func mainPage(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/main"{
    http.NotFound(w, r)
    return
  }
  s1, _ := template.ParseFiles("templates/base.html", "templates/main.html")
  s1.Execute(w, nil)
}
