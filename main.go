package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Struct per gestire i dati di tracciamento
type TimeTracking struct {
	SocialTime int // Tempo passato sui social
	WebTime    int // Tempo passato su altre pagine web
}

// Template HTML
var tmpl = template.Must(template.ParseFiles("index.html"))

// Funzione handler per la home page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Dati iniziali
	tracking := TimeTracking{
		SocialTime: 0,
		WebTime:    0,
	}

	// Rendering del template con i dati di tracciamento
	tmpl.Execute(w, tracking)
}

func main() {
	http.HandleFunc("/", homeHandler)
	// Servire file statici (per il banner)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server in esecuzione su http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

import  (
    "gopkg.in/gomail.v2"
	
)

func sendEmail(reportContent string) {
    m := gomail.NewMessage()
    m.SetHeader("From", "emanuele.zuffranieri@gmail.com")
    m.SetHeader("To", "emanuele.zuffranieri@gmail.com")
    m.SetHeader("Subject", "Report del tempo trascorso")
    m.SetBody("text/plain", reportContent)

    d := gomail.NewDialer("smtp.gmail.com", 587, "emanoele.zuffranieri@gmail.com", "Pesciolina8183@")

    if err := d.DialAndSend(m); err != nil {
        fmt.Println("Errore invio email:", err)
    }
}