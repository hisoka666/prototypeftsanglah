package main

import (
   "log"
   "net/http"
   "html/template"
)

func main(){
http.HandleFunc("/", index)
http.HandleFunc("/main", mainPage)
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

func mainPage(w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/main"{
    http.NotFound(w, r)
    return
  }
  s1, _ := template.ParseFiles("templates/base.html", "templates/main.html")
  s1.Execute(w, nil)
}
