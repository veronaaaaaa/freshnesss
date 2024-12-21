package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/services", servicesPage)
	http.HandleFunc("/schedule", schedulePage)
	http.HandleFunc("/contacts", contactsPage)
	http.HandleFunc("/registration", registrationPage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func servicesPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "services.html")
}

func schedulePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "schedule.html")
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "contacts.html")
}

func registrationPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "registration.html")
}
