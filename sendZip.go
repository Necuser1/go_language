package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/getZipFile", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applicaiton/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=sample-zip-file.zip")

		http.ServeFile(w, r, "sample-zip-file.zip")
	})
	fmt.Println("starting server at 8080")
	http.ListenAndServe(":8080", nil)
}
