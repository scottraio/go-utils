package utils

import (
	"log"
	"math"
	"os"

	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
)

func HasError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Info(msg string) {
	println(color.Ize(color.Green, msg))
}

func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	radlat1 := float64(math.Pi * lat1 / 180)
	radlat2 := float64(math.Pi * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(math.Pi * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / math.Pi
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}

// use godot package to load/read the .env file and
// return the value of the key
func GetDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		log.Fatal("Could not load .env file")
	}

	return os.Getenv(key)
}
