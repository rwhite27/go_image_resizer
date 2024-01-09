package main

import (
	"image"
	"os"
)

func main() {

	file, err := os.Open("./astro.png")

	if err != nil {
		println(err)
	}

	defer file.Close()

	_, imageString, err := image.Decode(file)

	if err != nil {
		println(err)
	}

	println(imageString)
}
