package main

import (
	"math/rand"
	"strconv"
	"time"
)

func imageSrc() string {
	randSize := rand.New(rand.NewSource(time.Now().UnixNano()))
	imageSrc := "https://source.unsplash.com/random/" + strconv.Itoa(randSize.Intn(1000)) + "x" + strconv.Itoa(randSize.Intn(1000))
	return imageSrc
}

func imageSrcByHeightWidth(minWidth int, maxWidth int, minHeight int, maxHeight int) string {
	rand.Seed(time.Now().Unix())
	imageSrc := "https://source.unsplash.com/random/" + strconv.Itoa(rand.Intn(maxWidth-minWidth)+minWidth) + "x" + strconv.Itoa(rand.Intn(maxHeight-minHeight)+minHeight)
	return imageSrc
}
