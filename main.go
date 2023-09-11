package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kalikidhar/A-MS-IN-GO/info"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Root Handler message")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "App is running OK")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("App health check")
	resp := map[string]string{
		"status":    "OK",
		"timestamp": time.Now().UTC().String(),
	}
	json.NewEncoder(w).Encode(resp)

}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get the IP info details")
	hostName, error := info.GetHostName()

	if error != nil {
		panic(error)
	}

	IP, error := info.GetIPAddress()

	fmt.Println(hostName, IP)

	resp := map[string]string{
		"hostname": hostName,
		"ip_addr":  IP.String(),
	}
	json.NewEncoder(w).Encode(resp)

	fmt.Println(hostName)

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/info", infoHandler)
	r.HandleFunc("/health", healthCheck)
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book titled %s with page %s", title, page)

	})
	log.Fatal(http.ListenAndServe(":80", r))
}
