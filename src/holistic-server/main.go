package main

import (
	"log"
	"net/http"
)

func ProcessRequest(w http.ResponseWriter, req *http.Request) {
	//w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(formatRequest(req)))

	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	buildHandler := http.FileServer(http.Dir("static"))
	http.Handle("/", buildHandler)

	http.HandleFunc("/api", ProcessRequest)

	err := http.ListenAndServeTLS(":5000", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
