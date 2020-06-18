package server

import (
	"math"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// get the next event for poisson distribution
// https://preshing.com/20111007/how-to-generate-random-timings-for-a-poisson-process/
// rate defined by a Poisson process with rps = λ
func nextTimePoisson(rps float64) float64 {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	return -math.Log(1-r.Float64()) / rps
}

// constant rate
func nextTimeLinear(rps float64) float64 {
	return 1.0 / rps
}

// SleepPoisson put the program in sleep state for a certain time
// sleep time is a poisson distribution with rps (rate per second)
// return the sleep time in micro seconds
func SleepPoisson(rps float64) int64 {
	nextTimeUs := int64(nextTimePoisson(rps / 1000000))
	time.Sleep(time.Microsecond * time.Duration(nextTimeUs))

	return nextTimeUs
}

// SleepLinear put the program in sleep state for a certain time
// sleep time is a linear distribution with rps (rate per second)
// return the sleep time in micro seconds
func SleepLinear(rps float64) int64 {
	nextTimeUs := int64(nextTimeLinear(rps / 1000000))
	time.Sleep(time.Microsecond * time.Duration(nextTimeUs))

	return nextTimeUs
}

func randomRune(length int) []rune {
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return b
}

// RandomString return a random string with length
func RandomString(length int) string {
	return string(randomRune(length))
}

// RandomByte returns a random byte array with length
func RandomByte(length int) []byte {
	return []byte(RandomString(length))
}
