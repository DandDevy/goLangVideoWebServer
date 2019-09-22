package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const videoFilePath string = "videos"

/*
HelloWorld to test if things are alright
 */
func sayHello(w http.ResponseWriter, r *http.Request) {

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	_, _ = w.Write([]byte(message))
}

func sayNamesOfVideos(w http.ResponseWriter, r *http.Request)  {

	//opens file
	f, err := os.Open(videoFilePath)
	if err != nil {
		log.Fatal(err)
	}

	videoFilesInfo, err := f.Readdir(-1)

	//closes file
	_ = f.Close()

	if err != nil {
		log.Fatal(err)
	}



	println("videoFilesInfo", videoFilesInfo)
	fmt.Printf("type:%T    size: %v\n", videoFilesInfo, len(videoFilesInfo))
	
	
	

	//sends the list of names
	for _, file := range videoFilesInfo {
		fmt.Println(file.Name())
		_, _ = w.Write([]byte(file.Name()))
	}

}

/*
Get the name of the first video.
 */
func sayNameOfFirstVideo(w http.ResponseWriter, r *http.Request)  {
	//message := "asd"

	//opens file
	f, err := os.Open(videoFilePath)
	if err != nil {
		log.Fatal(err)
	}

	files, err := f.Readdir(-1)

	//closes file
	_ = f.Close()

	if err != nil {
		log.Fatal(err)
	}

	//sends the list of names
	for _, file := range files {
		fmt.Println(file.Name())
		_, _ = w.Write([]byte(file.Name()))
	}

	//_, _ = w.Write([]byte(message))
}



func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/getFirstVideoName", sayNameOfFirstVideo)
	http.HandleFunc("/getVideoFileNames", sayNamesOfVideos)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}