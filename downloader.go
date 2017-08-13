package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func download(totalDownload int, isRandom bool, minWidth int, maxWidth int, minHeight int, maxHeight int) {
	i := 1
	for i <= totalDownload {
		if isRandom == true {
			// Download a random image
			downloadStatus := downloadFile(imageSrc())
			downloadMessage(downloadStatus)
		} else {
			// Download images with a defined height and width.
			downloadStatus := downloadFile(imageSrcByHeightWidth(minWidth, maxWidth, minHeight, maxHeight))
			downloadMessage(downloadStatus)
		}
		i = i + 1
	}
}

func downloadMessage(downloadStatus bool) {
	if downloadStatus == true {
		fmt.Println("✔ Downloaded a image")
	} else {
		fmt.Println("✘ Failed to downloaded a image")
	}
}

func randomFileName(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func downloadFile(url string) bool {

	// url := "http://i.imgur.com/m1UIjW1.jpg"

	// Get the image
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed to download this image.")
	}

	defer response.Body.Close() // Finaly close the body res.

	// Create the images directory
	os.Mkdir("./images", os.FileMode(0777))

	// Open a file for writing.
	file, err := os.Create("./images/" + randomFileName(25) + ".jpg")
	if err != nil {
		fmt.Println("Unable to write on disk.")
	}

	// Dump the data to a file.
	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println("Unable to save image on disk.")
	}

	// Close the open file
	file.Close()

	return true

}
