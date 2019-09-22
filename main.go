package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const videoFilePath string = "videos"

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/getVideoNames", sayVideoNames)
	http.HandleFunc("/getVideoFileDetails", giveVideoFileDetails)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

/*
HelloWorld to test if things are alright
 */
func sayHello(w http.ResponseWriter, r *http.Request) {

	println("sayHello")
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	_, _ = w.Write([]byte(message))
}

type VideoDetail struct {
	Name string
	Size int
}

func giveVideoFileDetails(w http.ResponseWriter, r *http.Request)  {
	println("giveVideoFileDetails")

	//set response to json
	w.Header().Set("Content-Type", "application/json")



	//opens file
	f, err := os.Open(videoFilePath)
	if err != nil {
		log.Fatal(err)
	}

	videoFilesInfo, err := f.Readdir(-1)

	//push close file
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	videoFilesInfoLength := len(videoFilesInfo)
	fmt.Printf("%v   %-T\n", videoFilesInfoLength, videoFilesInfoLength)

	var (
		videoDetails [2]VideoDetail
	)

	for i, _ := range videoFilesInfo{
		println("videoFilesInfo[i].Name() : ",videoFilesInfo[i].Name(), "videoFilesInfo[i].Size() : ", videoFilesInfo[i].Size())
	}
	videoDetails[0] = VideoDetail{
		Name: "video.mp4",
		Size: 4}

	fmt.Printf("%t", videoDetails)

	json.NewEncoder(w).Encode(videoDetails)

	//println("videoFilesInfo", videoFilesInfo)
	//fmt.Printf("type:%T    size: %v\n", videoFilesInfo, len(videoFilesInfo))


}

/*
Get the name of the first video.
 */
func sayVideoNames(w http.ResponseWriter, r *http.Request)  {
	println("sayVideoNames")
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
