package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	sf "github.com/trmiller/safe_payments"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	
	details := sf.PaymentDetails{
		SourceAccount:      r.FormValue("sourceAccount"),
		DestinationAccount: r.FormValue("destinationAccount"),
		Amount:             r.FormValue("amount"),
	}

	// do something with details
	_ = details

	tmpl.Execute(w, struct{ Success bool }{true})
}

func payment(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	pd  := sf.PaymentDetails{
		SourceAccount:      r.FormValue("sourceAccount"),
		DestinationAccount: r.FormValue("destinationAccount"),
		Amount:             r.FormValue("amount"),
	}

	sf.Pay(&pd)

	//thanks(r,w)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/payment", payment)
	http.HandleFunc("/thanks", thanks)
}

func thanks(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("thanks.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	tmpl.Execute(w, struct{ Success bool }{true})
}

func main() {

	setupRoutes()

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	fmt.Println("Go Web App Started on Port " + httpPort)
	http.ListenAndServe(":"+httpPort, nil)
}
