package sample

import (
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomImageName() string {
	min := 0
	max := 10000
	randomNumber := min + rand.Int()%(max-min+1)
	return "img" + strconv.Itoa(randomNumber)
}
