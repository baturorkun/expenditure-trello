package main

import (
	"expenditure/lib"
	"expenditure/setting"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)


type PageData struct {
	Title 	string
	Msg 	string
}


func cleanup() {
	fmt.Println("Cleanup")
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	setting.Setup()
	lib.Setup()

	// mux serves
	mx := mux.NewRouter()
	mx.HandleFunc("/", Enter)
	mx.HandleFunc("/save", Save)


	// http serves
	http.Handle("/", mx)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe( setting.ServerSetting.Port, nil)
}


func Enter(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		return
	}

	log.Println("Request ->", r.URL.Path)

	tmpl := template.Must(template.ParseFiles("templates/layouts/default.html", "templates/expenditure-form.html",))

	r.ParseForm()

	tmpl.Execute(w, PageData{Title: setting.AppSetting.Title, Msg: r.Form.Get("msg") })

}


func Save(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/save" {
		return
	}

	log.Println("Request ->", r.URL.Path)

	r.ParseForm()


	card := lib.CreateCard(r.PostForm)

	//lib.AddCard(card)

	http.Redirect(w, r, "/?msg=Your expenditure card was saved. " + card.Name, http.StatusPermanentRedirect)
}

