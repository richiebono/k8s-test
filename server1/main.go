package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Full Cycle Esquenta server 1</h1>"))
	})
	http.ListenAndServe(":8081", nil)
}
