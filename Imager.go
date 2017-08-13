package main

import (
	"fmt"
	"strings"
)

func main() {

	// Display the banner.
	intro()

	// Determine how many images need to download from the user.
	howManyImagesToDownload := takeIntInput("How many images to download?")

	if howManyImagesToDownload > 0 {

		// Determine if randomly generate width and height of each downloaded image.
		downloadRandom := strings.ToLower(takeStringInput("Download images in random width x height? y/n"))

		if downloadRandom == "y" || downloadRandom == "yes" {

			// Randomly generate image width and height.
			fmt.Print("Downloading images ...")
			download(howManyImagesToDownload, true, 0, 0, 0, 0)

		} else {

			// Download images in a given height and width length.

			// Take minimum width value
			minWidth := takeIntInput("Minimum image width:")

			// Take maximum width value
			maxWidth := takeIntInput("Maximum image width:")

			// Take minimum height value
			minHeight := takeIntInput("Minimum image height:")

			// Take maximum height value
			maxHeight := takeIntInput("Maximum image height:")

			download(howManyImagesToDownload, false, minWidth, maxWidth, minHeight, maxHeight)

		}
	} else {
		fmt.Println("Nothing to donwload, Bye.")
	}

	fmt.Println()
	fmt.Println()
}
