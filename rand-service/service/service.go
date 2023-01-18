package service

import (
	"math/rand"
	"time"
)

func GetRand(max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max)
}
