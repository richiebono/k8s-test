package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Full Cycle Esquenta server2</h1>"))
	})
	http.ListenAndServe(":8082", nil)
}
