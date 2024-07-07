package main

import "net/http"

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /fetch", Fetch)

	http.ListenAndServe(":3001", router)
}

func Fetch(w http.ResponseWriter, r *http.Request) {

}
