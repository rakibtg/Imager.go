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
