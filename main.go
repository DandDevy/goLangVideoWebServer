package main

import (
	"net/http"
	"strings"
)

/*
HelloWorld to test if things are alright
 */
func sayHello(w http.ResponseWriter, r *http.Request) {

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	_, _ = w.Write([]byte(message))
}

/*
Get the name of the first video.
 */
func sayNameOfFirstVideo(w http.ResponseWriter, r *http.Request)  {
	message := "asd"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/getFirstVideoName", sayNameOfFirstVideo)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}