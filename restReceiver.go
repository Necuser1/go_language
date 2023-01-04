package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	var unmarshalData interface{}
	err = json.Unmarshal(data, &unmarshalData)
	fmt.Println(unmarshalData)
	fmt.Println(err)
	w.WriteHeader(200)

}
func main() {
	http.HandleFunc("/ngsi-ld/listen", Handler)
	http.ListenAndServe(":8080", nil)
}

